package data

import (
	fish "gitlab.com/james.bondoc.dev/simple-fish-api/app/model"
)

func GetFishes() *[]fish.FishModel {

	fishList := []fish.FishModel{
		fish.FishModel{
			Id:          1,
			Image:       "https://images.immediate.co.uk/production/volatile/sites/4/2020/03/GettyImages-1058304880-c-0b54061-scaled.jpg?quality=90&resize=768%2C574",
			Name:        "Axolotl",
			Description: "Axolotls are members of the tiger salamander, or Ambystoma tigrinum, species complex, along with all other Mexican species of Ambystoma.",
			Habitat:     []string{"River", "Lake", "Waterfall"},
			Price:       23,
		},
		fish.FishModel{
			Id:          2,
			Image:       "https://www.aquariumofpacific.org/images/made_new/images-uploads-19_Queensland_Grouper_-_Robin_Riggs_600_q85.jpg",
			Name:        "Grouper",
			Description: "Groupers are fish of any of a number of genera in the subfamily Epinephelinae of the family Serranidae, in the order Perciformes.",
			Habitat:     []string{"Sea"},
			Price:       100,
		},
		fish.FishModel{
			Id:          3,
			Image:       "https://www2.illinois.gov/dnr/education/PublishingImages/WAFalligatorGar.jpg",
			Name:        "Alligator Gar",
			Description: "The alligator gar is a ray-finned euryhaline fish related to the bowfin in the infraclass Holostei. It is the biggest species in the gar family, and among the largest freshwater fish in North America.",
			Habitat:     []string{"River", "Lake"},
			Price:       70,
		},
		fish.FishModel{
			Id:          4,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/a/ae/Katri.jpg",
			Name:        "Gold Fish",
			Description: "The goldfish is a freshwater fish in the family Cyprinidae of order Cypriniformes. It is commonly kept as a pet in indoor aquariums",
			Habitat:     []string{"Pond"},
			Price:       10,
		},
		fish.FishModel{
			Id:          5,
			Image:       "https://www.aquariumnexus.com/wp-content/uploads/2020/06/redtail-catfish.jpg",
			Name:        "Red tail catfish",
			Description: "The redtail catfish, Phractocephalus hemioliopterus, is a pimelodid catfish. In Venezuela, it is known as cajaro, and in Brazil, it is known as pirarara, stemming from the Tupi language words pirá and arara. It is the only extant species of the genus Phractocephalus.",
			Habitat:     []string{"River", "Lake"},
			Price:       25,
		},
		fish.FishModel{
			Id:          6,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Pygocentrus_nattereri_-_Karlsruhe_Zoo_01.jpg/1200px-Pygocentrus_nattereri_-_Karlsruhe_Zoo_01.jpg",
			Name:        "Piranha",
			Description: "A piranha or piraña, a member of family Serrasalmidae, or a member of the subfamily Serrasalminae within the tetra family, Characidae in order Characiformes",
			Habitat:     []string{"River", "Lake"},
			Price:       30,
		},
		fish.FishModel{
			Id:          7,
			Image:       "https://exoticfishshop.net/wp-content/uploads/2019/12/exoticfishshop.com-red-wolf-fish-01.jpg",
			Name:        "Common Wolf fish",
			Description: "Hoplias malabaricus, also known as the wolf fish, tiger fish, guabine or trahira, is a predatory Central and South American freshwater ray finned fish of the characiform family Erythrinidae.",
			Habitat:     []string{"River", "Lake"},
			Price:       80,
		},
	}

	return &fishList
}
