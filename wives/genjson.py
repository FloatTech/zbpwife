import os
import json

sgs = {}
if os.path.exists('wife.json'):
    with open('wife.json', 'r',encoding='utf-8') as f:
        sgs = json.load(f)
# print(sgs)
pwd = os.getcwd()
urlList = []
for i in sgs:
    urlList.append(sgs[i]['url'])
ind = len(sgs)
print(os.listdir(pwd))
for i in os.listdir(pwd):
    if i.split('.')[1] == "jpg" or i.split('.')[1] == "png":
        if i in urlList:
            continue
        print(i)
        sgs[ind] = {"name":i.split('.')[0],'url':i,"lines":[]}
        ind+=1
with open('wife.json', 'w+',encoding='utf-8') as f:
    json.dump(sgs, f,ensure_ascii=False)