<?xml version="1.0" encoding="UTF-8"?>
<!--
Copyright 2019 Philippe Martin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:fo="http://www.w3.org/1999/XSL/Format"
                version="1.0">
  <xsl:import href="/usr/share/xml/docbook/stylesheet/docbook-xsl/fo/docbook.xsl"/>

  <xsl:param name="paper.type">A4</xsl:param>
  <xsl:param name="double.sided">1</xsl:param>
  <xsl:param name="variablelist.as.blocks">1</xsl:param>
  <xsl:param name="section.autolabel">1</xsl:param>
  <xsl:param name="generate.toc">
    book      toc,title
    part      title
  </xsl:param>
  <xsl:attribute-set name="monospace.verbatim.properties">
    <xsl:attribute name="wrap-option">wrap</xsl:attribute>
  </xsl:attribute-set>
  <xsl:param name="insert.link.page.number">yes</xsl:param>
  <xsl:param name="index.on.type">1</xsl:param>
</xsl:stylesheet>
