import random
import string


random_strings = lambda : "".join([random.choice(list(string.ascii_letters)) for _ in range(random.randint(10, 100))])

def jscript(file_name):
	with open(file_name, 'r') as handle:
		context = handle.read()


	for code_line in context.split(";"): # random varibale name
		if not code_line.find("var") == -1 and not code_line.find("=") == -1:
			index_start = code_line.find("var") + 3 
			index_end 	= code_line.find("=")
			variable 	= code_line[index_start:index_end].strip()
			variableTxt = code_line[index_end+1:].strip()
			if variableTxt[0] == '"' and variableTxt[-1] == '"' or variableTxt[0] == "'" and variableTxt[-1] == "'":
				split_variableTxt = list(variableTxt[1:-1])
				split_variableTxt = "'" + "'+'".join(split_variableTxt) + "'"
				context = context.replace(variableTxt, split_variableTxt)
			context = context.replace(variable, random_strings())


	for code_line in context.splitlines(): # ranom name functions and arguments
		if not code_line.find("function") == -1 and not code_line.find("){") == -1:
			iteam = code_line[code_line.find("function")+8:code_line.find("{")].strip()
			args = iteam[iteam.find("(")+1:iteam.rfind(")")].split(",")
			function = iteam[:iteam.find("(")].strip()
			context = context.replace(function, random_strings())
			for arg in args:
				if arg:
					context = context.replace(arg, random_strings())



	# end obfuscation
	context = context.replace(';\n', ';')
	context = context.replace('\n', '')
	#print(context)
	return context

