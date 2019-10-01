# coding: UTF-8
from bs4 import BeautifulSoup
from selenium import webdriver
import chromedriver_binary
from selenium.webdriver.chrome.options import Options

# ブラウザのオプションを格納する変数
options = Options()

# Headlessモードを有効にする
options.set_headless(True)

# ブラウザを起動する
driver = webdriver.Chrome(chrome_options=options)

# ブラウザでアクセスする
driver.get("https://techbookfest.org/event/tbf07/circle")

# HTMLを文字コードをUTF-8に変換してから取得
html = driver.page_source.encode('utf-8')

# BeautifulSoupで扱えるようにパース
soup = BeautifulSoup(html, "html.parser")

# ファイル出力
file_text = open("url_data.txt", "w")
elems = soup.find_all("a")
for elem in elems:
    print(elem.get('href'),file=file_text)
file_text.close()

