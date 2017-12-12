package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func parseToHTML(issueBody, token string) (string, error) {
	body, err := json.Marshal(map[string]string{"text": issueBody})
	if err != nil {
		return "", err
	}

	html, err := post("https://api.github.com/markdown", token, body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}

func saveFile(filename, html string) error {
	os.Mkdir("articles", os.ModeDir)
	html = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

<!-- Begin Jekyll SEO tag v2.3.0 -->
<title>chyroc.github.io | 基于issues的个人博客</title>
<meta property="og:title" content="chyroc.github.io" />
<meta property="og:locale" content="en_US" />
<meta name="description" content="基于issues的个人博客" />
<meta property="og:description" content="基于issues的个人博客" />
<link rel="canonical" href="http://blog.chyroc.cn/articles/some.html" />
<meta property="og:url" content="http://blog.chyroc.cn/articles/some.html" />
<meta property="og:site_name" content="chyroc.github.io" />
<script type="application/ld+json">
{"name":null,"description":"基于issues的个人博客","author":null,"@type":"WebPage","url":"http://blog.chyroc.cn/articles/some.html","image":null,"publisher":null,"headline":"chyroc.github.io","dateModified":null,"datePublished":null,"sameAs":null,"mainEntityOfPage":null,"@context":"http://schema.org"}</script>
<!-- End Jekyll SEO tag -->

    <link href="http://blog.chyroc.cn/assets/css/style.css?v=efb62de7b7525b6b14733e28358121421ac29cfd" rel="stylesheet">
  </head>
  <body>
    <div class="container-lg px-3 my-5 markdown-body">

        <h1><a href="http://blog.chyroc.cn/">chyroc.github.io</a></h1>

` + html + `    <div id="disqus_thread"></div>
<script>

/**
*  RECOMMENDED CONFIGURATION VARIABLES: EDIT AND UNCOMMENT THE SECTION BELOW TO INSERT DYNAMIC VALUES FROM YOUR PLATFORM OR CMS.
*  LEARN WHY DEFINING THESE VARIABLES IS IMPORTANT: https://disqus.com/admin/universalcode/#configuration-variables*/
/*
var disqus_config = function () {
this.page.url = PAGE_URL;  // Replace PAGE_URL with your page's canonical URL variable
this.page.identifier = PAGE_IDENTIFIER; // Replace PAGE_IDENTIFIER with your page's unique identifier variable
};
*/
(function() { // DON'T EDIT BELOW THIS LINE
var d = document, s = d.createElement('script');
s.src = 'https://chyroc.disqus.com/embed.js';
s.setAttribute('data-timestamp', +new Date());
(d.head || d.body).appendChild(s);
})();
</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>



<script id="dsq-count-scr" src="//chyroc.disqus.com/count.js" async></script>




  </body>`
	return ioutil.WriteFile(filename, []byte(html), 0644)
}