package validators

import (
	"sort"
	"strings"

	"github.com/lowl11/super-utils/util"
)

var ListFullFormToAbr = []AbrToFullForm{
	{AbrName: "АО", FullName: "Акционерное общество"},
	{AbrName: "АОЗТ", FullName: "Акционерное общество закрытого типа"},
	{AbrName: "ГККП", FullName: "Государственное коммунальное казенное предприятие"},
	{AbrName: "ГКП", FullName: "Государственное коммунальное предприятие"},
	{AbrName: "ГКП на ПХВ", FullName: "Государственное коммунальное предприятие на праве хозяйственного ведения"},
	{AbrName: "ГКП на ПХВ", FullName: "Государственное коммунальное предприятие на правах хозяйственного ведения"},
	{AbrName: "ГКП на ПХВ", FullName: "КГП на ПХВ"},
	{AbrName: "ГСПК", FullName: "Гаражно-строительный потребительский кооператив"},
	{AbrName: "ЖПК", FullName: "Жилищный потребительский кооператив"},
	{AbrName: "ЖСПК", FullName: "Жилищно-строительный потребительский кооператив"},
	{AbrName: "ЗАО", FullName: "Закрытое акционерное общество"},
	{AbrName: "ИП", FullName: "Индивидуальный предприниматель"},
	{AbrName: "КГУ", FullName: "Коммунальное государственное учреждение"},
	{AbrName: "КТ", FullName: "Коммандитное товарищество"},
	{AbrName: "КФХ", FullName: "Крестьянское (фермерское) хозяйство"},
	{AbrName: "КХ", FullName: "Крестьянское хозяйство"},
	{AbrName: "КЭАС", FullName: "Кооператив по эксплуатации автомобильных стоянок"},
	{AbrName: "НАО", FullName: "Непубличное акционерное общество"},
	{AbrName: "НИИ", FullName: "Научно-исследовательский институт"},
	{AbrName: "НКО", FullName: "Некоммерческая организация"},
	{AbrName: "НПП", FullName: "Научно-производственное предприятие"},
	{AbrName: "НПЦ", FullName: "Научно-производственный центр"},
	{AbrName: "ОАО", FullName: "Открытое акционерное общество"},
	{AbrName: "ОДО", FullName: "Общество с дополнительной ответственностью"},
	{AbrName: "ООО", FullName: "Общество с ограниченной ответственностью"},
	{AbrName: "ПАО", FullName: "Публичное акционерное общество"},
	{AbrName: "ПК", FullName: "Производственный кооператив"},
	{AbrName: "ПКГ", FullName: "Потребительский кооператив по газификации"},
	{AbrName: "ПТ", FullName: "Полное товарищество"},
	{AbrName: "ТДО", FullName: "Товарищество с дополнительной ответственностью"},
	{AbrName: "РГП", FullName: "Республиканское государственное предприятие"},
	{AbrName: "РГП на ПХВ", FullName: "Республиканское государственное предприятие на праве хозяйственного ведения"},
	{AbrName: "РГП на ПХВ", FullName: "Республиканское государственное предприятие на правах хозяйственного ведения"},
	{AbrName: "СТ", FullName: "Садоводческое товарищество"},
	{AbrName: "ТОО", FullName: "Товарищество с ограниченной ответственностью"},
	{AbrName: "ТС", FullName: "Товарищество собственников"},
	{AbrName: "ФХ", FullName: "Фермерское хозяйство"},
	{AbrName: "AG", FullName: "Aktiengesellschaft"},
	{AbrName: "BV", FullName: "Vennootschap Met Beperkte Aansparkelij kheid"},
	{AbrName: "Corp.", FullName: "Corporation"},
	{AbrName: "GmbH", FullName: "Gesellschaft mit beschrakter Haftung"},
	{AbrName: "IBC", FullName: "International Business Company"},
	{AbrName: "IC", FullName: "International Company"},
	{AbrName: "Inc.", FullName: "Incorporated"},
	{AbrName: "LDC", FullName: "Limited Duration Company"},
	{AbrName: "LP", FullName: "Limited Partnership"},
	{AbrName: "LLC", FullName: "Limited Liability Company"},
	{AbrName: "LLC", FullName: "Компания с ограниченной ответственностью по обязательствам"},
	{AbrName: "LLC", FullName: "Общество с ограниченной ответственностью по обязательствам"},
	{AbrName: "LLC", FullName: "Товарищество с ограниченной ответственностью по обязательствам"},

	{AbrName: "Ltd", FullName: "Limited"},
	{AbrName: "NV", FullName: "Naamlose Vennootschap"},
	{AbrName: "PLC", FullName: "Public Limited Company"},

	{AbrName: "SA", FullName: "Sosiedad Anonima"},
	{AbrName: "SA", FullName: "Societe Anonyme"},

	{AbrName: "SARL", FullName: "Societe a Responsidilite Limitee"},
}

func FullNameToAbbreviation(companyName string) (string, bool) {
	valid := IsValidCompanyName(companyName)
	if !valid {
		return companyName, false
	}

	return transformFullNameToShortName(companyName), true
}

func transformFullNameToShortName(companyName string) string {
	companyName = util.RemoveExtraSpaces(companyName)

	// Сортировка по длине аббревиатуры
	// При сравнении исключаем кейс "ГКП" преобразуется, "ГКП на ПХВ" нет
	abrTransformList := ListAbrToFullForm
	sort.Slice(abrTransformList, func(i, j int) bool {
		return len([]rune(abrTransformList[i].FullName)) > len([]rune(abrTransformList[j].FullName))
	})

	for _, transform := range abrTransformList {
		if len([]rune(companyName)) < len([]rune(transform.FullName))+1 {
			continue
		}
		//в конце полного имени компании должен быть пробел
		fullName := strings.ToLower(transform.FullName) + " "
		arr := []rune(companyName)

		if strings.EqualFold(fullName, string(arr[0:len([]rune(fullName))])) {
			bef := string(arr[0 : len([]rune(fullName))-1])
			return strings.Replace(companyName, bef, transform.AbrName, 1)
		}
	}
	return companyName
}
