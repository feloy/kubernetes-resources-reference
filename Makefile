default:
	@echo "commands: clean, docbook, pdf"

clean:
	rm -rf build

docbook: clean build/k8s-api.xml

build/k8s-api.xml: $(wildcard *.go **/*.go) config.yaml
	mkdir -p build
	go run main.go > build/k8s-api.xml

FORMAT ?= USletter
pdf: build/k8s-api.xml
	(cd build && \
	mkdir -p pdf-$(FORMAT) && \
	cd pdf-$(FORMAT) && \
	xsltproc --stringparam fop1.extensions 1 --stringparam paper.type $(FORMAT) -o k8s-api-$(FORMAT).fo ../../xsl/api.xsl ../k8s-api.xml && \
	fop -pdf k8s-api-$(FORMAT).pdf -fo k8s-api-$(FORMAT).fo && \
	rm  k8s-api-$(FORMAT).fo)

pdf-6x9in: build/k8s-api.xml
	(cd build && \
	mkdir -p pdf && \
	cd pdf && \
	xsltproc --stringparam fop1.extensions 1 -o k8s-api.fo ../../xsl/api-6x9in.xsl ../k8s-api.xml && \
	fop -pdf k8s-api.pdf -fo k8s-api.fo && \
	rm  k8s-api.fo)

test:
	@echo $(FORMAT)
