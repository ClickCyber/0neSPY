import sys, base64
import os, random
import string



create_variable = lambda : "".join([random.choice(list(string.ascii_letters)) for _ in range(random.randint(5, 50))])


def obfuscation(file_name):
	with open(file_name, 'r') as handle:
		context = handle.read()


	for code_line in context.split(";"):
		if not code_line.find("var") == -1 and not code_line.find("=") == -1:
			index_start = code_line.find("var") + 3 
			index_end 	= code_line.find("=")
			variable 	= code_line[index_start:index_end].strip()
			variableTxt = code_line[index_end+1:].strip()
			if variableTxt[0] == '"' and variableTxt[-1] == '"' or variableTxt[0] == "'" and variableTxt[-1] == "'":
				split_variableTxt = list(variableTxt[1:-1])
				split_variableTxt = "'" + "'+'".join(split_variableTxt) + "'"
				context = context.replace(variableTxt, split_variableTxt)
			context = context.replace(variable, create_variable())
			print(variable, create_variable())
	# end obfuscation
	context = context.replace(';\n', ';')
	print(context)
	return context


if len(sys.argv) < 1:
	quit("Usage:python3 CrypterJScript.py <file name>")
elif os.path.exists(sys.argv[1]):
	obfuscation(sys.argv[1])
else:
	quit("[-] File Is Not Exists")
