package validators

import (
	"regexp"

	"github.com/lowl11/super-utils/util"
)

var validCharsForCorp = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9ӘәІіҢңҒғҮүҰұҚқӨөҺһЁё"' \-,./«»\\–—-]+$`)

var OrgForms = map[string]string{
	"АО":         "Акционерное общество",
	"АОЗТ":       "Акционерное общество закрытого типа",
	"ГККП":       "Государственное коммунальное казенное предприятие",
	"ГКП":        "Государственное коммунальное предприятие",
	"ГКП на ПХВ": "Государственное коммунальное предприятие на праве хозяйственного ведения",
	"ГСПК":       "Гаражно-строительный потребительский кооператив",
	"ЖПК":        "Жилищный потребительский кооператив",
	"ЖСПК":       "Жилищно-строительный потребительский кооператив",
	"ЗАО":        "Закрытое акционерное общество",
	"ИП":         "Индивидуальный предприниматель",
	"КГУ":        "Коммунальное государственное учреждение",
	"КТ":         "Коммандитное товарищество",
	"КФХ":        "Крестьянское (фермерское) хозяйство",
	"КХ":         "Крестьянское хозяйство",
	"КЭАС":       "Кооператив по эксплуатации автомобильных стоянок",
	"НАО":        "Непубличное акционерное общество",
	"НИИ":        "Научно-исследовательский институт",
	"НКО":        "Некоммерческая организация",
	"НПП":        "Научно-производственное предприятие",
	"НПЦ":        "Научно-производственный центр",
	"ОАО":        "Открытое акционерное общество",
	"ОДО":        "Общество с дополнительной ответственностью",
	"ООО":        "Общество с ограниченной ответственностью",
	"ПАО":        "Публичное акционерное общество",
	"ПК":         "Производственный кооператив",
	"ПКГ":        "Потребительский кооператив по газификации",
	"ПТ":         "Полное товарищество",
	"ТДО":        "Товарищество с дополнительной ответственностью",
	"РГП":        "Республиканское государственное предприятие",
	"РГП на ПХВ": "Республиканское государственное предприятие на праве хозяйственного ведения",
	"СТ":         "Садоводческое товарищество",
	"ТОО":        "Товарищество с ограниченной ответственностью",
	"ТС":         "Товарищество собственников",
	"ФХ":         "Фермерское хозяйство",
	"AG":         "Aktiengesellschaft",
	"BV":         "Vennootschap Met Beperkte Aansparkelij kheid",
	"Corp.":      "Corporation",
	"GmbH":       "Gesellschaft mit beschrakter Haftung",
	"IBC":        "International Business Company",
	"IC":         "International Company",
	"Inc.":       "Incorporated",
	"LDC":        "Limited Duration Company",
	"LLC":        "Limited Liability Company",
	"Ltd":        "Limited",
	"NV":         "Naamlose Vennootschap",
	"PLC":        "Public Limited Company",
	"SA":         "Sosiedad Anonima",
	"SARL":       "Societe a Responsidilite Limitee",
}

func CompanyShortNameTransform(companyName string) string {
	var fullName string
	re := regexp.MustCompile(`^(.*?)\s*«(.*?)»$`)
	match := re.FindStringSubmatch(companyName)
	if len(match) != 3 {
		return companyName
	}

	if orgForm, found := util.GetKeyByValue(OrgForms, match[1]); found {
		// Найдено совпадение, создайте полное наименование
		fullName = orgForm + " " + "«" + match[2] + "»"
		companyName = fullName
	}

	return companyName
}

func IsValidCompanyName(companyName string) bool {
	// Проверка на допустимые символы
	if !validCharsForCorp.MatchString(companyName) {
		return false
	}
	return true
}
