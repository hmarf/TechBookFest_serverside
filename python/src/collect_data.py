# coding: UTF-8
from bs4 import BeautifulSoup
from selenium import webdriver
import chromedriver_binary
from selenium.webdriver.chrome.options import Options
import pickle
import sys

# データを保存する
i = 0
sys.setrecursionlimit(10000)
with open('../data/getData.txt', 'wb') as ff:

    # 保存するデータ
    save_datas = []

    # ブラウザのオプションを格納する変数
    options = Options()

    # Headlessモードを有効にする
    options.set_headless(True)

    # ブラウザを起動する
    driver = webdriver.Chrome(chrome_options=options)

    urlHeader = "https://techbookfest.org"
    pathfile = "../data/url_data.txt"

    with open(pathfile) as f:
        for _path in f:
            i += 1
            url = urlHeader + _path
            print(i,url)
            # ブラウザでアクセスする
            driver.get(url)

            # HTMLを文字コードをUTF-8に変換してから取得
            html = driver.page_source.encode('utf-8')

            # BeautifulSoupで扱えるようにパース
            soup = BeautifulSoup(html, "html.parser")

            circle = soup.find(class_="circle-name").string
            arrange = soup.find(class_="circle-space-label").string
            genre = soup.find(class_="circle-genre-label").string
            keyword = soup.find(class_="circle-genre-free-format").string

            circle_image = soup.find(class_="circle-detail-image").find(class_="ng-star-inserted")

            book_title = []
            for a in soup.find_all(class_="mat-card-title"):
                book_title.append(a.string)

            book_content = []
            for a in soup.find_all(class_="products-description"):
                book_content.append(a.string)

            for title, content in zip(book_title, book_content):
                data = [circle,circle_image['src'],arrange,genre,keyword,title,content,url]
                save_datas.append(data)

    pickle.dump(save_datas,ff)
