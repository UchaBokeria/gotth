import html
import json
import os
import sys
import time
import requests
from bs4 import BeautifulSoup

BaseUrl = "https://yacco.com"
# InternalUploadUrl = "http://localhost:3000/upload/"

Categories = [
    "automobile",
    "moto-quad-karting",
    "transport-t-p",
    "agriculture-motoculture",
    "boites-et-ponts",
    "entretien-et-nettoyage",
    "marine-nautisme",
    "aviation-legere",
    "specialites",
]

Products = []
GoProducts = '''package main

import (
    uploader "main/server/common/helpers"
	"main/server/common/globals"
	"main/server/common/storage"
	"main/server/model"
)

var AllProductsSeed = []model.Products{
'''

opening = "{"
closing = "}"
for CategoryID, Category in enumerate(Categories):
    page = 0
    while True:
        page += 1
        print("Request endpoint: " + BaseUrl + "/fr/nos-produits/" + Category + "/" + str(page))
        CategoryRawHtml = requests.get(BaseUrl + "/fr/nos-produits/" + Category + "/" + str(page)).text
        CategoryHtml = BeautifulSoup(CategoryRawHtml, "html.parser")
        ProductContainers = CategoryHtml.select("#results > div > div > a")
        
        if(len(ProductContainers) < 1): break

        for ProductContainer in ProductContainers:
            DetailUrl = ProductContainer.get("href")
            print("Request endpoint: " + BaseUrl + DetailUrl)
            ProductRawHtml = requests.get(BaseUrl + DetailUrl).text
            ProductHtml = BeautifulSoup(ProductRawHtml, "html.parser")
            print("Product " + DetailUrl.replace("https://yacco.com/fr/nos-produits/automobile/","") + "Is Found And Fetched.")

            Title = ProductHtml.select_one("body > div.main-wrapper > main > section.section-product-view > div > div.row.mb-5 > div.col-lg-5 > div > h2")

            DescriptionHtml = ProductHtml.select_one("body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-7.mt-5.mt-lg-0 > div")
            Description = ""
            DescriptionParagraphs = ProductHtml.select("body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-7.mt-5.mt-lg-0 > div > p")
            for DescriptionParagraph in DescriptionParagraphs:
                Description += DescriptionParagraph.get_text()

            Packing = []
            GoPackings = ""
            PackingRows = ProductHtml.select("body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-5 > div:nth-child(3) > table > tbody > tr")
            for index, PackingRow in enumerate(PackingRows):
                Name = ProductHtml.select_one(        "body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-5 > div:nth-child(3) > table > tbody > tr:nth-child(" + str(index + 1) + ") > td:nth-child(1) span")
                Reference = ProductHtml.select_one(   "body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-5 > div:nth-child(3) > table > tbody > tr:nth-child(" + str(index + 1) + ") > td:nth-child(2)")
                Conditioning = ProductHtml.select_one("body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-5 > div:nth-child(3) > table > tbody > tr:nth-child(" + str(index + 1) + ") > td:nth-child(3)")
                Cardboard = ProductHtml.select_one(   "body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-5 > div:nth-child(3) > table > tbody > tr:nth-child(" + str(index + 1) + ") > td:nth-child(4)")
                
                Pack = {
                    "Name": getattr(Name, 'get_text', lambda: "")() or "",
                    "Reference": getattr(Reference, 'get_text', lambda: "")() or "",
                    "Conditioning":getattr(Conditioning, 'get_text', lambda: "")() or "",
                    "Cardboard": getattr(Cardboard, 'get_text', lambda: "")() or "",
                }

                Packing.append(Pack)
                GoPackings += f''' {opening}Name: "{Pack.get("Name")}", Slug: "{Pack.get("Reference")}", Reference: "{Pack.get("Conditioning")}", Conditioning: "{Pack.get("Cardboard")}", Cardboard: "{Pack.get("Name")}"{closing}, \n '''

            Specifications = []
            GoSpecifications = ""
            SpecificationsRows = ProductHtml.select("body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-5 > div:nth-child(1) > div > p")
            for SpecificationsRow in (SpecificationsRows):
                specName = getattr(SpecificationsRow, 'get_text', lambda: "")()
                Specification = {
                    "Name": specName,
                    "Slug": getattr(specName, 'replace', lambda: "")(" ", "_")
                }
                Specifications.append(Specification)
                GoSpecifications += f''' {opening}Name: "{Specification.get("Name")}", Slug: "{Specification.get("Slug")}"{closing}, \n '''

                
            Approvals = []
            GoApprovals = ""
            ApprovalsRows = ProductHtml.select("body > div.main-wrapper > main > section.section-product-view > div > div:nth-child(3) > div.col-lg-5 > div:nth-child(1) > div > p")
            for ApprovalsRow in (ApprovalsRows):
                ApprovalsName = getattr(ApprovalsRow, 'get_text', lambda: "")()
                Approval = {
                    "Name": ApprovalsName,
                    "Slug": getattr(ApprovalsName, 'replace', lambda: "")(" ", "_")
                }
                Approvals.append(Approval)
                GoApprovals += f''' {opening}Name: "{Approval.get("Name")}", Slug: "{Approval.get("Slug")}"{closing}, \n '''

            Properties = []
            GoProperties = ""
            PropertiesRows = ProductHtml.select("body > div.main-wrapper > main > section.section-product-view > div > div.row.mb-5 > div.col-lg-5 > div > ul > li > span")
            for PropertiesRow in (PropertiesRows):
                PropertiesName = getattr(PropertiesRow, 'get_text', lambda: "")()
                Propertie = {
                    "Name": PropertiesName,
                    "Slug": getattr(PropertiesName, 'replace', lambda: "")(" ", "_")
                }
                Properties.append(Propertie)
                GoProperties += f''' {opening}Name: `{Propertie.get("Name")}`, Slug: `{Propertie.get("Slug")}` {closing}, \n '''

            TechnialSheet = ProductHtml.select_one("body > div.main-wrapper > main > section.section-product-view > div > div.row.mb-5 > div.col-lg-5 > div > div > a:nth-child(2)")
            TechnialSheetUrl = TechnialSheet.get("href")

            ThumbnailHtml = ProductHtml.select_one("body > div.main-wrapper > main > section.section-product-view > div > div.row.mb-5 > div.col-lg-7 > picture > img")
            response = requests.get(BaseUrl + ThumbnailHtml.get("src"))
            Thumbnail = '''{opening} {closing}'''

            # Check if the request was successful (status code 200)
            if response.status_code != 200:
                print("Product " + Title.get_text() + "Failed To Download " + ThumbnailHtml.get("src"), response.status_code)
            else:
                basename, extension = os.path.splitext(ThumbnailHtml.get("src"))
                parts = basename.split("/")
                fileName = parts[-1] + extension

                with open("./public/uploads/products/" + fileName, "wb") as file:
                    file.write(response.content)
                
                with open("./public/uploads/products/" + fileName, 'rb') as file:
                    Thumbnail = f''' model.Files{opening}
                        Name:       `{fileName}`,
                        Original:   `{fileName}`,
                        Location:   "local",
                        Path:       "/uploads/{fileName}",
                        Size:       {file.__sizeof__()},
                        Base64:     "",
                        Compressed: false,
                        TypeID:     uploader.GetDbTypeIdByExtension(`{extension.replace(".", "")}`),
                    {closing}, \n '''
                    # files = {'file': file}
                    # Send the file using a POST request
                    # ImgUploadResponse = requests.post(InternalUploadUrl, files=files)

            # ThumbnailID = int(ImgUploadResponse.json()["ID"])

            Product = {
                "Name": html.escape(Title.get_text()),
                "Slug": html.escape(Title.get_text().replace(" ", "_")),
                "Description": html.escape(Description),
                "DescriptionHtml": html.escape(DescriptionHtml.prettify()),
                "CategoryID": CategoryID + 1,
                "TechnialSheetUrl": TechnialSheetUrl,
                "Thumbnail": Thumbnail,
                "Packing": Packing,
                "Approvals": Approvals,
                "Properties": Properties,
                "Specifications": Specifications,
            }

            Products.append(Product)
            print("Prodcut " + Product.get("Name") + " Is Parsed.")


            GoProducts += f'''{opening}
            Name: `{Product.get("Name")}`,
            Slug: `{Product.get("Slug")}`,
            Description: `{Product.get("Description")}`,
            DescriptionHtml: `{Product.get("DescriptionHtml")}`,
            CategoryID: {Product.get("CategoryID")},
            TechnicalSheetUrl: "{Product.get("TechnialSheetUrl")}",
            Thumbnail: {Product.get("Thumbnail")},
            Packing: []model.Product_packaging{opening}
                {(GoPackings)}
            {closing},
            Approvals: []model.Product_approvals{opening}
                {(GoSpecifications)}
            {closing},
            Properties: []model.Product_properties{opening}
                {(GoApprovals)}
            {closing},
            Specifications: []model.Product_specifications{opening}
                {(GoProperties)}
            {closing},
        {closing},
            '''
            time.sleep(3)

GoProducts += '''
}

func main() {
	globals.SetupEnvironmentVariables()
	storage.Connect(storage.Default())
    for _, row := range AllProductsSeed { storage.DB.Create(&row) }
}
'''

# cleaning
GoProducts.replace("™", "")


with open("./cmd/parser/main.go", "w", encoding='utf-8') as file:
    file.write(GoProducts)

with open("./cmd/parser/result.json", "w") as json_file:
    json.dump(Products, json_file, indent=4)
