package exporters

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
	"brief/briefExporter/common"
)

const (
	titleGroup          = "title"
	originalTitleGroup  = "original_title"
	authorGroup         = "author"
	pageGroup           = "page"
	firstLocationGroup  = "location"
	secondLocationGroup = "second_location"
	createdOnDateGroup  = "created_date"
	createdOnTimeGroup  = "created_time"
	noteDataGroup       = "note_data"
	recordTypeGroup     = "record_type"
)

type (
	KindleExporter struct {}

	noteData struct {
		titleNoteData map[string]string
		noteData      map[string]string
	}
)

var recordTypesToSkip = []string{"Highlight", "Bookmark"}

func (m *KindleExporter) GetNotes(path string) ([]*NoteRecord, error) {

	log.Printf("Processing file %s:\n", path)

	fileData, err := common.GetFileData(path)
	common.Check(err)

	go func() {
		common.GetFileChecksum(fileData)
	}()

	str := string(fileData)

	recordRegexp := regexp.MustCompile(`(g?)(i?)(?P<` + titleGroup + `>[\wА-Яа-яіїєґ'#\-*:*\s*\.*\,*]+)\s` +
		`(?P<` + originalTitleGroup + `>\({1}[\wА-Яа-яіїєґ\s*\.*\,*]+\){1})?\s?` +
		`(\({1}(?P<` + authorGroup + `>[\wА-Яа-яіїєґ\;*\s*\.*\,*]+)\){1}){1}` +
		`[\r\n]*-\sYour\s(?P<` + recordTypeGroup + `>(Note|Highlight|Bookmark))\son\s` +
		`(page\s(?P<` + pageGroup + `>[\d]+)\s\|\s)?` +
		`Location\s(?P<` + firstLocationGroup + `>[\d]+)\-?` +
		`(?P<` + secondLocationGroup + `>[\d]+)?\s\|\sAdded\son\s` +
		`[\w]+\,{1}\s(?P<` + createdOnDateGroup + `>[\w]+\s[\d]+\,\s\d{4})` +
		`\s(?P<` + createdOnTimeGroup + `>\d{1,2}:\d{2}:\d{2}\s(AM|PM))` +
		`[\r\n]*(?P<` + noteDataGroup + `>[\S\s]+)[\r\n]*`)

	split := regexp.MustCompile("={10}[\r\n]*").Split(str, -1)

	var notes []*NoteRecord

	i := 0
	for i < len(split)-1 {
		var noteData noteData
		titleGroup := common.GetGroupsData(recordRegexp, split[i])

		//handling of Highlights and Bookmarks
		if common.Contains(recordTypesToSkip, titleGroup[recordTypeGroup]) {
			i++
			continue
		}

		noteData.titleNoteData = titleGroup
		noteData.noteData = common.GetGroupsData(recordRegexp, split[i+1])
		note := &NoteRecord{}

		checkNoteFiled(noteData, note.checkTitle, note.checkAltTitle,
			note.checkAuthor, note.checkPage, note.checkLocations, note.checkDateAndTime, note.checkNoteTitleAndText)
		notes = append(notes, note)

		i += 2
	}
	return notes, nil
}

func checkNoteFiled(data noteData, fns ...func(data noteData) (*NoteRecord, error)) (err error) {
	for _, fn := range fns {
		if _, err = fn(data); err != nil {
			break
		}
	}
	return
}

func (note *NoteRecord) checkTitle(data noteData) (*NoteRecord, error) {
	if baseNoteFieldCheck(data, titleGroup, false) {
		note.BookTitle = data.titleNoteData[titleGroup]
		return note, nil
	}
	return note, errors.New(titleGroup + " could not be processed further")
}

func (note *NoteRecord) checkAltTitle(data noteData) (*NoteRecord, error) {
	if baseNoteFieldCheck(data, originalTitleGroup, true) {
		note.BookOriginalName = data.titleNoteData[originalTitleGroup]
		return note, nil
	}
	return note, errors.New(originalTitleGroup + " could not be processed further")
}

func (note *NoteRecord) checkAuthor(data noteData) (*NoteRecord, error) {
	if baseNoteFieldCheck(data, authorGroup, false) {
		authors := strings.Split(data.titleNoteData[authorGroup], ";")
		if len(authors) > 1 {
			for i := range authors {
				parsedAuthor := regexp.MustCompile(",\\\\s*").Split(authors[i], -1)
				if len(parsedAuthor) > 1 {
					common.Reverse(parsedAuthor)

					author, err := handleAuthor(parsedAuthor)
					if err != nil {
						return note, errors.New(authorGroup + " could not be processed further")
					}
					note.BookAuthor = append(note.BookAuthor, *author)

				} else {
					authorData := strings.Split(data.titleNoteData[authorGroup], " ")

					author, err := handleAuthor(authorData)
					if err != nil {
						return note, errors.New(authorGroup + " could not be processed further")
					}
					note.BookAuthor = append(note.BookAuthor, *author)
				}
			}
		} else {
			parsedAuthor := regexp.MustCompile(",\\\\s*").Split(data.titleNoteData[authorGroup], -1)
			if len(parsedAuthor) > 1 {
				common.Reverse(parsedAuthor)

				author, err := handleAuthor(parsedAuthor)
				if err != nil {
					return note, errors.New(authorGroup + " could not be processed further")
				}
				note.BookAuthor = append(note.BookAuthor, *author)
			} else {
				authorData := strings.Split(data.titleNoteData[authorGroup], " ")

				author, err := handleAuthor(authorData)
				if err != nil {
					return note, errors.New(authorGroup + " could not be processed further")
				}

				note.BookAuthor = append(note.BookAuthor, *author)
			}
		}
		return note, nil
	}
	return note, errors.New(authorGroup + " could not be processed further")
}

func (note *NoteRecord) checkPage(data noteData) (*NoteRecord, error) {
	var err error
	if data.noteData[pageGroup] != "" || data.titleNoteData[pageGroup] != "" {
		if data.titleNoteData[pageGroup] != "" {
			note.FirstPage, err = strconv.Atoi(data.titleNoteData[pageGroup])
		}
		if data.noteData[pageGroup] != "" {
			note.SecondPage, err = strconv.Atoi(data.noteData[pageGroup])
		}
		return note, err
	}
	return note, nil
}

func (note *NoteRecord) checkLocations(data noteData) (*NoteRecord, error) {
	var err error
	if data.noteData[firstLocationGroup] != "" || data.titleNoteData[firstLocationGroup] != "" {
		if data.titleNoteData[firstLocationGroup] != "" {
			note.FirstLocation, err = strconv.Atoi(data.titleNoteData[firstLocationGroup])
		}
		if data.noteData[secondLocationGroup] != "" {
			note.SecondLocation, err = strconv.Atoi(data.noteData[secondLocationGroup])
		}
		return note, err
	}
	return note, errors.New(firstLocationGroup + " could not be processed further")
}

func (note *NoteRecord) checkNoteTitleAndText(data noteData) (*NoteRecord, error) {
	if data.noteData[noteDataGroup] != "" {
		note.NoteTitle = data.titleNoteData[noteDataGroup]
		note.NoteText = data.noteData[noteDataGroup]
		return note, nil
	}
	return note, errors.New(noteDataGroup + " could not be processed further")
}

func (note *NoteRecord) checkDateAndTime(data noteData) (*NoteRecord, error) {
	if baseNoteFieldCheck(data, createdOnDateGroup, false) || baseNoteFieldCheck(data, createdOnTimeGroup, false) {
		note.CreatedOn = data.noteData[createdOnDateGroup] + " " + data.noteData[createdOnTimeGroup]
		return note, nil
	}
	return note, errors.New(noteDataGroup + " could not be processed further")
}

func baseNoteFieldCheck(data noteData, groupName string, isOptional bool) bool {
	if !isOptional {
		if data.titleNoteData[groupName] != "" && data.titleNoteData[groupName] == data.noteData[groupName] {
			return true
		}
	} else {
		return true
	}
	return false
}

func handleAuthor(authorData []string) (*Author, error) {
	author := &Author{}
	switch len(authorData) {
	case 1:
		author.FirstName = authorData[0]
	case 2:
		author.FirstName = authorData[0]
		author.Surname = authorData[1]
	case 3:
		author.FirstName = authorData[0]
		author.SecondaryName = authorData[1]
		author.Surname = authorData[2]
	default:
		return &Author{FirstName: "Authors unknown"}, nil
	}
	return author, nil
}
