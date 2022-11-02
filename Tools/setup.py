import sys
import os

if not len(sys.argv) > 1:
	quit("Usage: python3 setup.py <action>")

def load_key():
	return "12345"

def install():
	pass

def build(ip, port):
	if not os.path.exists('./Tools/tmp/'):
		os.mkdir('./Tools/tmp/')
	with open('./Client/main.go') as handle:
		source = handle.read()
	source = source.replace('{IP}', ip)
	source = source.replace('"{PORT}"', port)
	source = source.replace('{KEY}', load_key())
	with open('./Tools/tmp/main.go', 'w+') as file:
		file.write(source)
	os.system("./Tools/complier.sh ./Tools/tmp/main.go")

action = sys.argv[1]

if action == "install":
	pass
elif action == "build":
	if len(sys.argv) > 3:
		ip = sys.argv[2]
		port = sys.argv[3]
		build(ip, port)
	else:
		quit("Usage: python3 build <ip> <port>")
