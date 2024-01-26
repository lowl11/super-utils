package validate

import (
	"regexp"
	"slices"
	"time"
)

const (
	DocumentTypePassport               = 1
	DocumentTypeCertificate            = 2
	DocumentTypeIdentification         = 3
	DocumentTypeResidence              = 4
	DocumentTypeOfStatelessPerson      = 5
	DocumentTypeDiplomaticIdentityCard = 6
	DocumentTypeForeignPassport        = 7
	DocumentTypeResidencyPermit        = 8
)

var docTypeList = []int{
	DocumentTypePassport, DocumentTypeCertificate, DocumentTypeIdentification,
	DocumentTypeResidence, DocumentTypeOfStatelessPerson, DocumentTypeDiplomaticIdentityCard,
	DocumentTypeForeignPassport, DocumentTypeResidencyPermit,
}

func DocumentType(documentType int) bool {
	return slices.Contains(docTypeList, documentType)
}

var (
	oneLetterAndOtherDigitsPattern = regexp.MustCompile(`^[A-Za-z]\d{8}$`)
	numberPattern                  = regexp.MustCompile(`^\d{9}$`)
	otherDocumentPattern           = regexp.MustCompile(`^[\p{L}\d\s\-.()/:\[\]{}\\]*$`)
)

func DocumentNumber(documentType int, documentNumber string) bool {
	switch documentType {
	case DocumentTypePassport:
		return oneLetterAndOtherDigitsPattern.MatchString(documentNumber)
	case DocumentTypeIdentification:
		return numberPattern.MatchString(documentNumber)
	case DocumentTypeResidence:
		return numberPattern.MatchString(documentNumber)
	case DocumentTypeOfStatelessPerson:
		return oneLetterAndOtherDigitsPattern.MatchString(documentNumber)
	case DocumentTypeCertificate, DocumentTypeForeignPassport,
		DocumentTypeDiplomaticIdentityCard, DocumentTypeResidencyPermit:
		// Проверка на наличие букв либо латиницы, либо кириллицы, но не одновременно
		return otherDocumentPattern.MatchString(documentNumber) && !(containsLatinAndCyrillic(documentNumber))
	default:
		return false // Неизвестный тип документа
	}
}

func DocumentIssueDate(format, value string) bool {
	parsed, err := time.Parse(format, value)
	if err != nil {
		return false
	}

	return !parsed.After(time.Now())
}

func DocumentExpireDate(format, value string, issueDate time.Time) bool {
	parsed, err := time.Parse(format, value)
	if err != nil {
		return false
	}

	return parsed.After(issueDate)
}

func containsLatinAndCyrillic(s string) bool {
	hasLatin := false
	hasCyrillic := false
	for _, char := range s {
		if 'A' <= char && char <= 'Z' || 'a' <= char && char <= 'z' {
			hasLatin = true
		}
		if 'А' <= char && char <= 'Я' || 'а' <= char && char <= 'я' {
			hasCyrillic = true
		}
	}
	return hasLatin && hasCyrillic
}
