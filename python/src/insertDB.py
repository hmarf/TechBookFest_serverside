# coding: UTF-8
import MySQLdb
import pickle
 
# データベースへの接続とカーソルの生成
connection = MySQLdb.connect(
    host='0.0.0.0',
    user='user',
    passwd='password',
    db='techBook')
cursor = connection.cursor()

# id, circle, circle_image, arr, genere, keyword, title, content
with open('../data/getData.txt','rb') as f:
    load_datas = pickle.load(f)
    for load_data in load_datas:
        data = []
        if load_data[6] == None:
            for dd in load_data:
                if dd == None:
                    continue
                dd = dd.replace('\'','')
                data.append(dd)
            sql = "INSERT INTO circle (circle, circle_image, arr, genere, keyword, title, circle_url) values ('" + data[0] + "','" + data[1] + "','" + data[2]+"','" + data[3]+"','" + data[4]+"','" + data[5]+"','" + data[6]+"')"
        else:
            for dd in load_data:
                dd = dd.replace('\'','')
                data.append(dd)
            sql = "INSERT INTO circle (circle, circle_image, arr, genere, keyword, title, content, circle_url) values ('" + data[0] + "','" + data[1] + "','" + data[2]+"','" + data[3]+"','" + data[4]+"','" + data[5]+"','" + data[6]+"','" + data[7]+"')"
            print(sql)
        cursor.execute(sql)

# 保存を実行
connection.commit()
 
# 接続を閉じる
connection.close()
