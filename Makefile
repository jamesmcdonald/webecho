all:
	make -C echoserver all
	make -C webtest all

clean:
	make -C echoserver clean
	make -C webtest clean
