package category

type CategoryFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FormatCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return formatter
}

func FormatCategories(categories []Category) []CategoryFormatter {
	if len(categories) == 0 {
		return []CategoryFormatter{}
	}

	var categoryFormatter []CategoryFormatter

	for _, category := range categories {
		formatter := FormatCategory(category)
		categoryFormatter = append(categoryFormatter, formatter)
	}

	return categoryFormatter
}
