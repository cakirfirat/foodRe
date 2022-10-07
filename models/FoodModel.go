package models

type RecognizedFood struct {
	FoodFamily []struct {
		ID   int     `json:"id"`
		Name string  `json:"name"`
		Prob float64 `json:"prob"`
	} `json:"foodFamily"`
	FoodType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"foodType"`
	ImageID       int `json:"imageId"`
	ModelVersions struct {
		Drinks       string `json:"drinks"`
		FoodType     string `json:"foodType"`
		Foodgroups   string `json:"foodgroups"`
		Foodrec      string `json:"foodrec"`
		Ingredients  string `json:"ingredients"`
		Ontology     string `json:"ontology"`
		Segmentation string `json:"segmentation"`
	} `json:"model_versions"`
	Occasion            string `json:"occasion"`
	SegmentationResults []struct {
		Center struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"center"`
		ContainedBbox struct {
			H int `json:"h"`
			W int `json:"w"`
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"contained_bbox"`
		FoodItemPosition   int `json:"food_item_position"`
		RecognitionResults []struct {
			FoodFamily []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"foodFamily"`
			FoodType struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"foodType"`
			ID         int     `json:"id"`
			Name       string  `json:"name"`
			Prob       float64 `json:"prob"`
			Subclasses []struct {
				FoodFamily []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"foodFamily"`
				FoodType struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"foodType"`
				ID   int     `json:"id"`
				Name string  `json:"name"`
				Prob float64 `json:"prob"`
			} `json:"subclasses"`
		} `json:"recognition_results"`
	} `json:"segmentation_results"`
}
