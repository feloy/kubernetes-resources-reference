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

  <!-- size specific -->
  <xsl:param name="page.height.portrait">9in</xsl:param>
  <xsl:param name="page.width.portrait">6in</xsl:param>
  <xsl:param name="page.margin.inner">0.75in</xsl:param>
  <xsl:param name="page.margin.outer">0.50in</xsl:param>
  <xsl:param name="page.margin.top">0.50in</xsl:param>
  <xsl:param name="page.margin.bottom">0.50in</xsl:param>
  <xsl:param name="body.font.master">8</xsl:param>
  <xsl:param name="body.start.indent">1pc</xsl:param>

  <xsl:param name="double.sided">1</xsl:param>
  <xsl:param name="variablelist.as.blocks">1</xsl:param>
  <xsl:param name="section.autolabel">1</xsl:param>
  <xsl:param name="generate.toc">
    book      toc,title
    part      title
  </xsl:param>
  <xsl:param name="toc.section.depth">1</xsl:param>
  <xsl:attribute-set name="monospace.verbatim.properties">
    <xsl:attribute name="wrap-option">wrap</xsl:attribute>
  </xsl:attribute-set>
  <xsl:param name="insert.link.page.number">yes</xsl:param>
  <xsl:param name="index.on.type">1</xsl:param>

  <!-- section margin -->
  <xsl:attribute-set name="section.title.properties">
    <xsl:attribute name="space-before.minimum">1.8em</xsl:attribute>
    <xsl:attribute name="space-before.optimum">2.0em</xsl:attribute>
    <xsl:attribute name="space-before.maximum">2.2em</xsl:attribute>
  </xsl:attribute-set>

  <!-- change chapter/appendix title format -->
  <xsl:template name="mychapter.title">
    <xsl:param name="node" select="."/>

    <fo:block xsl:use-attribute-sets="chap.label.properties">
      <xsl:call-template name="gentext">
        <xsl:with-param name="key">
          <xsl:choose>
            <xsl:when test="$node/self::chapter">chapter</xsl:when>
            <xsl:when test="$node/self::appendix">appendix</xsl:when>
          </xsl:choose>
        </xsl:with-param>
      </xsl:call-template>
      <xsl:text> </xsl:text>
      <xsl:apply-templates select="$node" mode="label.markup"/>
    </fo:block>
    <fo:block xsl:use-attribute-sets="chap.title.properties">
      <xsl:apply-templates select="$node" mode="title.markup"/>
    </fo:block>
  </xsl:template>

  <xsl:template match="title" mode="chapter.titlepage.recto.auto.mode">
    <fo:block xsl:use-attribute-sets="chapter.titlepage.recto.style" font-size="24.8832pt" font-weight="bold">
      <xsl:call-template name="mychapter.title">
        <xsl:with-param name="node" select="ancestor-or-self::chapter[1]"/>
      </xsl:call-template>
    </fo:block>
  </xsl:template>

  <xsl:template match="title" mode="appendix.titlepage.recto.auto.mode">
    <fo:block xsl:use-attribute-sets="appendix.titlepage.recto.style" margin-left="{$title.margin.left}" font-size="24.8832pt" font-weight="bold" font-family="{$title.fontset}">
      <xsl:call-template name="mychapter.title">
        <xsl:with-param name="node" select="ancestor-or-self::appendix[1]"/>
      </xsl:call-template>
    </fo:block>
  </xsl:template>
    
  <xsl:attribute-set name="chap.label.properties">
    <xsl:attribute name="font-size">12pt</xsl:attribute>
    <xsl:attribute name="text-align">right</xsl:attribute>
  </xsl:attribute-set>
  <xsl:attribute-set name="chap.title.properties">
    <xsl:attribute name="text-align">right</xsl:attribute>
    <xsl:attribute name="margin-bottom">1in</xsl:attribute>
  </xsl:attribute-set>

  <!-- change part title format -->
  <xsl:template name="mypart.title">

    <xsl:param name="node" select="."/>

    <fo:block xsl:use-attribute-sets="part.label.properties">
      <xsl:call-template name="gentext">
        <xsl:with-param name="key">
          <xsl:choose>
            <xsl:when test="$node/self::part">part</xsl:when>
          </xsl:choose>
        </xsl:with-param>
      </xsl:call-template>
      <xsl:text> </xsl:text>
      <xsl:apply-templates select="$node" mode="label.markup"/>
    </fo:block>
    <fo:block xsl:use-attribute-sets="part.title.properties">
      <xsl:apply-templates select="$node" mode="title.markup"/>
    </fo:block>
  </xsl:template>

  <xsl:template match="title" mode="part.titlepage.recto.auto.mode">
    <fo:block xsl:use-attribute-sets="part.titlepage.recto.style" text-align="center" font-size="24.8832pt" space-before="18.6624pt" font-weight="bold" font-family="{$title.fontset}">
      <xsl:call-template name="mypart.title">
        <xsl:with-param name="node" select="ancestor-or-self::part[1]"/>
      </xsl:call-template>
    </fo:block>
  </xsl:template>

  <xsl:attribute-set name="part.label.properties">
    <xsl:attribute name="font-size">12pt</xsl:attribute>
    <xsl:attribute name="text-align">right</xsl:attribute>
  </xsl:attribute-set>
  <xsl:attribute-set name="part.title.properties">
    <xsl:attribute name="text-align">right</xsl:attribute>
  </xsl:attribute-set>

  <!-- change header content -->

  <xsl:template name="header.content">
    <xsl:param name="pageclass" select="''"/>
    <xsl:param name="sequence" select="''"/>
    <xsl:param name="position" select="''"/>
    <xsl:param name="gentext-key" select="''"/>

    <fo:block>
      <!-- sequence can be odd, even, first, blank -->
      <!-- position can be left, center, right -->
      <xsl:choose>
        <xsl:when test="$sequence = 'blank'">
          <!-- nothing -->
        </xsl:when>
  
        <xsl:when test="$position='left'">
          <!-- Same for odd, even, empty, and blank sequences -->
          <xsl:call-template name="draft.text"/>
        </xsl:when>
  
        <xsl:when test="($sequence='odd') and $position='center'">
          <xsl:if test="$pageclass != 'titlepage'">
            <xsl:choose>
              <xsl:when test="ancestor::book and ($double.sided != 0)">
                <fo:retrieve-marker retrieve-class-name="section.head.marker"
                                    retrieve-position="first-including-carryover"
                                    retrieve-boundary="page-sequence"/>
              </xsl:when>
              <xsl:otherwise>
                <xsl:apply-templates select="." mode="titleabbrev.markup"/>
              </xsl:otherwise>
            </xsl:choose>
          </xsl:if>
        </xsl:when>

        <xsl:when test="($sequence='even') and $position='center'">
          <xsl:if test="$pageclass != 'titlepage'">
            <xsl:choose>
              <xsl:when test="ancestor::book and ($double.sided != 0)">
                <xsl:apply-templates select="." mode="titleabbrev.markup"/>
              </xsl:when>
              <xsl:otherwise>
                <xsl:apply-templates select="." mode="titleabbrev.markup"/>
              </xsl:otherwise>
            </xsl:choose>
          </xsl:if>
        </xsl:when>

        <xsl:when test="$position='center'">
          <!-- nothing for empty and blank sequences -->
        </xsl:when>
  
        <xsl:when test="$position='right'">
          <!-- Same for odd, even, empty, and blank sequences -->
          <xsl:call-template name="draft.text"/>
        </xsl:when>
  
        <xsl:when test="$sequence = 'first'">
          <!-- nothing for first pages -->
        </xsl:when>
  
        <xsl:when test="$sequence = 'blank'">
          <!-- nothing for blank pages -->
        </xsl:when>
      </xsl:choose>
    </fo:block>
  </xsl:template>
</xsl:stylesheet>
