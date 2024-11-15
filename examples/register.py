#!/usr/bin/python3
import json
import subprocess as proc

name = __file__.removesuffix('.py')
js = ''
with proc.Popen(name+".uidl", stdin=proc.PIPE, stdout=proc.PIPE, close_fds=True) as p:
    js = p.stdout.read()
ret = p.wait()

if ret == 0:
    data = json.loads(js)
    print('data:', data)
    print(f"""INSERT INTO person (name, email, gender, bio)
      VALUES ('{data["name"]}', '{data["email"]}', {data["gender"]}, '{data["bio"]}')""")
print("ret:", ret)
