package main

import (
    uploader "main/server/common/helpers"
	"main/server/common/globals"
	"main/server/common/storage"
	"main/server/model"
)

var AllProductsSeed = []model.Products{

    {
		Name: `AERO AVX 1000 4T 10W40`,
		Slug: `AERO_AVX_1000_4T_10W40`,
		Description: `L’huile Yacco AERO AVX 1000 4T 10W40 est une huile de technologie de synthèse pour moteurs 4 temps de l&#x27;aviation légère et ultralégère.L’huile Yacco AERO AVX 1000 4T 10W40 est une huile de technologie de synthèse développée afin de répondre aux dernières exigences des différents motoristes de l’aviation légère et ultralégère. Elle est particulièrement adaptée aux moteurs 4 Temps (Carburateur ou injection) atmosphériques ou turbocompressés, y compris ceux utilisés dans des conditions d’utilisation très sévères.L’huile Yacco AERO AVX 1000 4T 10W40 dispose d’une formulation exclusive permettant de répondre également aux exigences des réducteurs intégrés et des limiteurs de couple.Propriétés et avantagesL’huile Yacco AERO AVX 1000 4T 10W40 procure une excellente stabilité à hautes températures et propriétés dispersives renforcées pour éviter la formation de boues et de vernis, afin de maintenir les moteurs dans un excellent état de propreté. Elle est adaptée et recommandée pour les moteurs fonctionnant au carburant - avec ou sans plomb - et essence aviation 100LL.L’huile Yacco AERO AVX 1000 4T 10W40 assure un bon comportement à froid pour faciliter les démarrages et augmenter la protection des moteurs par une lubrification plus rapide. Elle offre également une protection renforcée contre l’usure et la corrosion.Ne pas utiliser cette huile dans les moteurs aviation où il est recommandé une huile sans cendre pour moteurs à pistons (Continental, Lycoming, Rolls-Royce, Jabiru, etc.`,
		DescriptionHtml: `&lt;div class=&quot;product-view-description rte&quot;&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AERO AVX 1000 4T 10W40
  &lt;/strong&gt;
  est une huile de technologie de synthèse pour moteurs 4 temps de l&#x27;aviation légère et ultralégère.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AERO AVX 1000 4T 10W40
  &lt;/strong&gt;
  est une huile de technologie de synthèse développée afin de répondre aux dernières exigences des différents motoristes de l’aviation légère et ultralégère. Elle
  &lt;strong&gt;
  &lt;/strong&gt;
  est particulièrement adaptée aux moteurs 4 Temps (Carburateur ou injection) atmosphériques ou turbocompressés, y compris ceux utilisés dans des conditions d’utilisation très sévères.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AERO AVX 1000 4T 10W40
  &lt;/strong&gt;
  dispose d’une formulation exclusive permettant de répondre également aux exigences des réducteurs intégrés et des limiteurs de couple.
 &lt;/p&gt;
 &lt;p&gt;
  &lt;strong&gt;
   Propriétés et avantages
  &lt;/strong&gt;
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AERO AVX 1000 4T 10W40
  &lt;/strong&gt;
  procure une excellente stabilité à hautes températures et propriétés dispersives renforcées pour éviter la formation de boues et de vernis, afin de maintenir les moteurs dans un excellent état de propreté. Elle est adaptée et recommandée pour les moteurs fonctionnant au carburant - avec ou sans plomb - et essence aviation 100LL.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AERO AVX 1000 4T 10W40
  &lt;/strong&gt;
  assure un bon comportement à froid pour faciliter les démarrages et augmenter la protection des moteurs par une lubrification plus rapide. Elle offre également une protection renforcée contre l’usure et la corrosion.
 &lt;/p&gt;
 &lt;p&gt;
  &lt;strong&gt;
   Ne pas utiliser cette huile dans les moteurs aviation où il est recommandé une huile sans cendre pour moteurs à pistons (Continental, Lycoming, Rolls-Royce, Jabiru, etc.
  &lt;/strong&gt;
 &lt;/p&gt;
&lt;/div&gt;
`,
		CategoryID: 1,
		TechnicalSheetUrl: "/upload/media/default/0001/16/37517e01f42a9a409c7921244565f97ceb5d0196.pdf",
		Thumbnail:  model.Files{
                        Name:       `thumb_807_product_big.jpeg`,
                        Original:   `thumb_807_product_big.jpeg`,
                        Location:   "local",
                        Path:       "/uploads/thumb_807_product_big.jpeg",
                        Size:       664,
                        Base64:     "",
                        Compressed: false,
                        TypeID:     uploader.GetDbTypeIdByExtension(`jpeg`),
                    },
		Packing: []model.Product_packaging{
			
		},
		Approvals: []model.Product_approvals{
			 {Name: "API SL", Slug: "API_SL"},  {Name: "JASO MA Performance", Slug: "JASO_MA_Performance"},  {Name: "Huile compatible avec les moteurs Rotax des séries 912 et 914 ainsi que les versions injectées 912i, 915i et 916i.", Slug: "Huile_compatible_avec_les_moteurs_Rotax_des_séries_912_et_914_ainsi_que_les_versions_injectées_912i,_915i_et_916i."}, 
		},
		Properties: []model.Product_properties{
			 {Name: "API SL", Slug: "API_SL"},  {Name: "JASO MA Performance", Slug: "JASO_MA_Performance"},  {Name: "Huile compatible avec les moteurs Rotax des séries 912 et 914 ainsi que les versions injectées 912i, 915i et 916i.", Slug: "Huile_compatible_avec_les_moteurs_Rotax_des_séries_912_et_914_ainsi_que_les_versions_injectées_912i,_915i_et_916i."}, 
		},
		Specifications: []model.Product_specifications{
			 {Name: `3359`, Slug: `3359` },  {Name: `Technologie de synthèse`, Slug: `Technologie_de_synthèse` },  {Name: `Viscosité :`, Slug: `Viscosité_:` }, 
		},
	},
            
    {
		Name: `AVX 1000 2T`,
		Slug: `AVX_1000_2T`,
		Description: `L’huile Yacco AVX 1000 2T est une huile 100% synthèse pour les moteurs 2 Temps de l’aviation légère.L’huile Yacco AVX 1000 2T est une huile 100% synthèse développée afin de répondre aux exigences des différents motoristes de l’aviation légère. Elle est particulièrement adaptée aux moteurs 2 Temps à injection ou à carburateur(s) y compris ceux utilisés dans des conditions très sévères (compétition, école, tractage de banderoles, etc.)Propriétés et avantagesL’huile Yacco AVX 1000 2T dispose de bases synthétiques esters apportant une résistance exceptionnelle aux hautes températures pour prévenir les risques de serrage.L’huile Yacco AVX 1000 2T réduit le calaminage des pots d’échappement, des lumières et des pistons. Elle convient aux graissages séparés ou aux mélanges.L’huile Yacco AVX 1000 2T procure d’excellentes performances antifumée, sa basse teneur en cendres empêchant la formation de dépôts dans la chambre de combustion. Son bas point d’écoulement permet une utilisation jusqu’à de très basses températures.`,
		DescriptionHtml: `&lt;div class=&quot;product-view-description rte&quot;&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 1000 2T
  &lt;/strong&gt;
  est une huile 100% synthèse pour les moteurs 2 Temps de l’aviation légère.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 1000 2T
  &lt;/strong&gt;
  est une huile 100% synthèse développée afin de répondre aux exigences des différents motoristes de l’aviation légère. Elle est particulièrement adaptée aux moteurs 2 Temps à injection ou à carburateur(s) y compris ceux utilisés dans des conditions très sévères (compétition, école, tractage de banderoles, etc.)
 &lt;/p&gt;
 &lt;p&gt;
  &lt;strong&gt;
   Propriétés et avantages
  &lt;/strong&gt;
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 1000 2T
  &lt;/strong&gt;
  dispose de bases synthétiques esters apportant une résistance exceptionnelle aux hautes températures pour prévenir les risques de serrage.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 1000 2T
  &lt;/strong&gt;
  réduit le calaminage des pots d’échappement, des lumières et des pistons. Elle convient aux graissages séparés ou aux mélanges.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 1000 2T
  &lt;/strong&gt;
  procure d’excellentes performances antifumée, sa basse teneur en cendres empêchant la formation de dépôts dans la chambre de combustion. Son bas point d’écoulement permet une utilisation jusqu’à de très basses températures.
 &lt;/p&gt;
&lt;/div&gt;
`,
		CategoryID: 1,
		TechnicalSheetUrl: "/upload/media/default/0001/04/2326dc1d851a57587516b721297004fbf2f6bbf2.pdf",
		Thumbnail:  model.Files{
                        Name:       `thumb_726_product_big.jpeg`,
                        Original:   `thumb_726_product_big.jpeg`,
                        Location:   "local",
                        Path:       "/uploads/thumb_726_product_big.jpeg",
                        Size:       664,
                        Base64:     "",
                        Compressed: false,
                        TypeID:     uploader.GetDbTypeIdByExtension(`jpeg`),
                    },
		Packing: []model.Product_packaging{
			
		},
		Approvals: []model.Product_approvals{
			 {Name: "API TC", Slug: "API_TC"},  {Name: "JASO FD", Slug: "JASO_FD"},  {Name: "ISO-L-EGD", Slug: "ISO-L-EGD"}, 
		},
		Properties: []model.Product_properties{
			 {Name: "API TC", Slug: "API_TC"},  {Name: "JASO FD", Slug: "JASO_FD"},  {Name: "ISO-L-EGD", Slug: "ISO-L-EGD"}, 
		},
		Specifications: []model.Product_specifications{
			 {Name: `3364`, Slug: `3364` },  {Name: `100% synthèse`, Slug: `100%_synthèse` }, 
		},
	},
            
    {
		Name: `AVX 500 4T 10W40 `,
		Slug: `AVX_500_4T_10W40_`,
		Description: `L’huile Yacco AVX 500 4T - SAE 10W40 est une huile 100% semi-synthèse développée afin de répondre aux exigences des différents motoristes de l’aviation légère.L’huile Yacco AVX 500 4T - SAE 10W40 est particulièrement adaptée aux moteurs atmosphériques 4 Temps turbocompressés, munis ou non de poussoirs hydrauliques y compris ceux utilisés dans des conditions de service sévères (école, tractage de banderoles, etc.).L’huile Yacco AVX 500 4T - SAE 10W40 dispose d’une formulation exclusive permettant de répondre également aux exigences des réducteurs intégrés et des limiteurs de couple.Propriétés et avantagesL’huile Yacco AVX 500 4T - SAE 10W40 procure une excellente stabilité de l’huile à hautes températures. Elle dispose de propriétés dispersives renforcées pour éviter la formation de dépôts.L’huile Yacco AVX 500 4T - SAE 10W40 est adaptée et recommandée pour les moteurs fonctionnant avec un carburant avec ou sans plomb.Ce produit est réservé à l’aviation légère. Il ne doit pas être utilisé pour les moteurs d’avions traditionnels tels Lycoming, Rolls-Royce, etc.`,
		DescriptionHtml: `&lt;div class=&quot;product-view-description rte&quot;&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 4T - SAE 10W40
  &lt;/strong&gt;
  est une huile 100% semi-synthèse développée afin de répondre aux exigences des différents motoristes de l’aviation légère.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 4T - SAE 10W40
  &lt;/strong&gt;
  est particulièrement adaptée aux moteurs atmosphériques 4 Temps turbocompressés, munis ou non de poussoirs hydrauliques y compris ceux utilisés dans des conditions de service sévères (école, tractage de banderoles, etc.).
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 4T - SAE 10W40
  &lt;/strong&gt;
  dispose d’une formulation exclusive permettant de répondre également aux exigences des réducteurs intégrés et des limiteurs de couple.
 &lt;/p&gt;
 &lt;p&gt;
  &lt;strong&gt;
   Propriétés et avantages
  &lt;/strong&gt;
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 4T - SAE 10W40
  &lt;/strong&gt;
  procure une excellente stabilité de l’huile à hautes températures. Elle dispose de propriétés dispersives renforcées pour éviter la formation de dépôts.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 4T - SAE 10W40
  &lt;/strong&gt;
  est adaptée et recommandée pour les moteurs fonctionnant avec un carburant avec ou sans plomb.
 &lt;/p&gt;
 &lt;p&gt;
  &lt;strong&gt;
   Ce produit est réservé à l’aviation légère.
  &lt;/strong&gt;
  Il ne doit pas être utilisé pour les moteurs d’avions traditionnels tels Lycoming, Rolls-Royce, etc.
 &lt;/p&gt;
&lt;/div&gt;
`,
		CategoryID: 1,
		TechnicalSheetUrl: "/upload/media/default/0001/16/9e5c05dfc1b7d4d72b85d38cef380781adee9366.pdf",
		Thumbnail:  model.Files{
                        Name:       `thumb_15935_product_big.jpg`,
                        Original:   `thumb_15935_product_big.jpg`,
                        Location:   "local",
                        Path:       "/uploads/thumb_15935_product_big.jpg",
                        Size:       664,
                        Base64:     "",
                        Compressed: false,
                        TypeID:     uploader.GetDbTypeIdByExtension(`jpg`),
                    },
		Packing: []model.Product_packaging{
			
		},
		Approvals: []model.Product_approvals{
			 {Name: "API SL", Slug: "API_SL"},  {Name: "Huile compatible pour les moteurs Rotax des séries 912 et 914.", Slug: "Huile_compatible_pour_les_moteurs_Rotax_des_séries_912_et_914."}, 
		},
		Properties: []model.Product_properties{
			 {Name: "API SL", Slug: "API_SL"},  {Name: "Huile compatible pour les moteurs Rotax des séries 912 et 914.", Slug: "Huile_compatible_pour_les_moteurs_Rotax_des_séries_912_et_914."}, 
		},
		Specifications: []model.Product_specifications{
			 {Name: `3360`, Slug: `3360` },  {Name: `Semi-synthèse`, Slug: `Semi-synthèse` },  {Name: `Viscosité :`, Slug: `Viscosité_:` }, 
		},
	},
            
    {
		Name: `AVX 500 2T`,
		Slug: `AVX_500_2T`,
		Description: `L’huile Yacco AVX 500 2T est une huile 100 % semi-synthèse pour moteurs 2 Temps de l’aviation légère.L’huile Yacco AVX 500 2T a été développée afin de répondre aux exigences des différents motoristes de l’aviation légère et ultralégère. Elle est particulièrement adaptée aux moteurs 2 Temps à injection ou à carburateur y compris ceux utilisés dans des conditions de service sévères (école, tractage de banderoles, etc.).Propriétés et avantagesL’huile Yacco AVX 500 2T dispose d’excellentes performances antifumées et une protection renforcée contre la formation de dépôt et le gommage des segments.L’huile Yacco AVX 500 2T est miscible instantanément dans le carburant, son bas point d’écoulement permettant une utilisation jusqu’à de très basses températures.Coloration de l’huile en bleu.`,
		DescriptionHtml: `&lt;div class=&quot;product-view-description rte&quot;&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 2T
  &lt;/strong&gt;
  est une huile 100 % semi-synthèse pour moteurs 2 Temps de l’aviation légère.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 2T
  &lt;/strong&gt;
  a été développée afin de répondre aux exigences des différents motoristes de l’aviation légère et ultralégère. Elle est particulièrement adaptée aux moteurs 2 Temps à injection ou à carburateur y compris ceux utilisés dans des conditions de service sévères (école, tractage de banderoles, etc.).
 &lt;/p&gt;
 &lt;p&gt;
  &lt;strong&gt;
   Propriétés et avantages
  &lt;/strong&gt;
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 2T
  &lt;/strong&gt;
  dispose d’excellentes performances antifumées et une protection renforcée contre la formation de dépôt et le gommage des segments.
 &lt;/p&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco AVX 500 2T
  &lt;/strong&gt;
  est miscible instantanément dans le carburant, son bas point d’écoulement permettant une utilisation jusqu’à de très basses températures.
 &lt;/p&gt;
 &lt;p&gt;
  Coloration de l’huile en bleu.
 &lt;/p&gt;
&lt;/div&gt;
`,
		CategoryID: 1,
		TechnicalSheetUrl: "/upload/media/default/0001/04/d70018fb4d9625fcfecf2c2c2f0c84d31388b458.pdf",
		Thumbnail:  model.Files{
                        Name:       `thumb_12089_product_big.jpeg`,
                        Original:   `thumb_12089_product_big.jpeg`,
                        Location:   "local",
                        Path:       "/uploads/thumb_12089_product_big.jpeg",
                        Size:       664,
                        Base64:     "",
                        Compressed: false,
                        TypeID:     uploader.GetDbTypeIdByExtension(`jpeg`),
                    },
		Packing: []model.Product_packaging{
			
		},
		Approvals: []model.Product_approvals{
			 {Name: "API TC+", Slug: "API_TC+"},  {Name: "Propriétés antifumées", Slug: "Propriétés_antifumées"}, 
		},
		Properties: []model.Product_properties{
			 {Name: "API TC+", Slug: "API_TC+"},  {Name: "Propriétés antifumées", Slug: "Propriétés_antifumées"}, 
		},
		Specifications: []model.Product_specifications{
			 {Name: `3365`, Slug: `3365` },  {Name: `Semi-synthèse`, Slug: `Semi-synthèse` }, 
		},
	},
            
    {
		Name: `Aé D80`,
		Slug: `Aé_D80`,
		Description: `L’huile Yacco Aé D80 est une huile minérale monograde dispersive pour moteurs à pistons des avions traditionnels tels Continental, Lycoming, Rolls-Royce, etc.`,
		DescriptionHtml: `&lt;div class=&quot;product-view-description rte&quot;&gt;
 &lt;p&gt;
  L’huile
  &lt;strong&gt;
   Yacco Aé D80
  &lt;/strong&gt;
  est une huile minérale monograde dispersive pour moteurs à pistons des avions traditionnels tels Continental, Lycoming, Rolls-Royce, etc.
 &lt;/p&gt;
&lt;/div&gt;
`,
		CategoryID: 1,
		TechnicalSheetUrl: "/upload/media/default/0001/01/ed28e74c1dca6d77efd936ce3b8f741e62dd8263.pdf",
		Thumbnail:  model.Files{
                        Name:       `thumb_536_product_big.jpeg`,
                        Original:   `thumb_536_product_big.jpeg`,
                        Location:   "local",
                        Path:       "/uploads/thumb_536_product_big.jpeg",
                        Size:       664,
                        Base64:     "",
                        Compressed: false,
                        TypeID:     uploader.GetDbTypeIdByExtension(`jpeg`),
                    },
		Packing: []model.Product_packaging{
			
		},
		Approvals: []model.Product_approvals{
			 {Name: "API SL", Slug: "API_SL"}, 
		},
		Properties: []model.Product_properties{
			 {Name: "API SL", Slug: "API_SL"}, 
		},
		Specifications: []model.Product_specifications{
			 {Name: `3366`, Slug: `3366` },  {Name: `Minérale`, Slug: `Minérale` },  {Name: `Viscosité :`, Slug: `Viscosité_:` }, 
		},
	},
            
}

func main() {
	globals.SetupEnvironmentVariables()
	storage.Connect(storage.Default())
    for _, row := range AllProductsSeed { storage.DB.Create(&row) }
}
