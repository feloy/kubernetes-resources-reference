# Copyright 2020 Philippe Martin
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

default:
	@echo "commands: clean, docbook, pdf, pdf-6x9in, epub, clean-website, website, serve"

clean:
	rm -rf build

docbook: clean build/k8s-api.xml

build/k8s-api.xml: $(wildcard *.go **/*.go) $(wildcard api/v1.19/* config/v1.19/*)
	mkdir -p build
	go run cmd/main.go docbook --config-dir config/v1.19/ --file api/v1.19/swagger.json > build/k8s-api.xml


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

epub: build/k8s-api.xml
	(cd build && \
	 rm -rf epub && mkdir -p epub && \
	 cd epub && \
	 xsltproc -o k8s.epub /usr/share/xml/docbook/stylesheet/docbook-xsl/epub3/chunk.xsl ../k8s-api.xml && \
	 zip -r k8s.epub META-INF mimetype OEBPS && \
	 rm -rf META-INF mimetype OEBPS)

clean-website:
	rm -rf website/content/en/docs/* website/public

website: clean-website
	mkdir -p website/content/en/docs
	go run cmd/main.go hugo --config-dir config/v1.19/ --file api/v1.19/swagger.json --output-dir website/content/en/docs
	(cd website && hugo)

clean-kwebsite:
	rm -rf kwebsite/content/en/docs/* kwebsite/public

kwebsite: clean-kwebsite
	mkdir -p kwebsite/content/en/docs
	go run cmd/main.go kwebsite --config-dir config/v1.19/ --file api/v1.19/swagger.json --output-dir kwebsite/content/en/docs

serve:
	(cd website/public && python3 -m http.server)
