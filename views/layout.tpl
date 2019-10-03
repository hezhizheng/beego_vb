<!DOCTYPE html>
<html lang="zh_cn">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="referrer" content="never">
  <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
  <meta name="description" content="茄哩啡">
  <meta name="author" content="茄哩啡">
  <title>茄哩啡</title>
  <link href="//cdn.bootcss.com/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  <link href="//cdn.bootcss.com/video.js/7.6.0/alt/video-js-cdn.min.css" rel="stylesheet">
  <!-- Custom styles for this template -->
  <link href="/static/css/blog.css" rel="stylesheet">
  <link href="//cdn.bootcss.com/github-markdown-css/3.0.1/github-markdown.min.css" rel="stylesheet">

</head>

<body>

{{/*nav*/}}
{{.nav}}

<div class="container">

  <div class="vb-content" >
    {{.content}}
  </div>

</div> <!-- /.container -->

{{.footer}}

<!-- Bootstrap core JavaScript
================================================== -->
<!-- jQuery (Bootstrap 的所有 JavaScript 插件都依赖 jQuery，所以必须放在前边) -->
<script src="//cdn.bootcss.com/jquery/3.4.1/jquery.min.js"></script>
<!-- 加载 Bootstrap 的所有 JavaScript 插件。你也可以根据需要只加载单个插件。 -->
<script src="//cdn.bootcss.com/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
{{/*<script src="//cdn.bootcss.com/video.js/7.6.0/alt/video.core.min.js"></script>*/}}

<script>
  $(function () { $("[data-toggle='tooltip']").tooltip(); });
</script>

<script>
  $(function (){

    // 上一页
    $(".vb-last-page").click(function(){
      const last_default_page = $("[last-default-page]").attr("last-default-page");

      const last_page_class = $(".vb-last-page").attr("class");

      if(last_page_class.indexOf("disabled") >= 0 ) {
        return false;
      }

      let w_url = window.location.href.split("/category/");
      let category_id = w_url[1] ? w_url[1] : '';
      let filter = {
        'category_id' : category_id,
        'p' : last_default_page
      };

      $.ajax({
        url:"/pagination",
        data:filter,
        success:function(result){
          // 启用下一页按钮
          $(".vb-next-page.disabled").attr("class","vb-next-page");
          $(".vb-next-page").attr("next-default-page",Number(result.PageNo) + 1);
          if ( !result.FirstPage )
          {
            $(".vb-last-page").attr("last-default-page",Number(result.PageNo) - 1);
          }else{
            // 最后一页禁用下一页
            $(".vb-last-page").attr("class","vb-last-page disabled");
          }

          let vb_history = '';
          $.each(result.List,function(index,value) {
            vb_history += "<div class=\"col-sm-6 col-md-3\" >\n" +
                    "                  <div class=\"thumbnail\">\n" +
                    "                  <a href=\"/article/" + value.Slug + "\">\n" +
                    "                  <div class=\"vb-thumb-img\"  style=\"width: 170px;\n" +
                    "          height: 100px;\n" +
                    "          min-width: 100%;\n" +
                    "          min-height: 100%;\n" +
                    "          background-repeat: no-repeat;\n" +
                    "          background-position: center center;\n" +
                    "          background-size: cover;\n" +
                    "          background-image:url("+ value.VideoCover + ")\">\n" +
                    "                  </div>\n" +
                    "                  </a>\n" +
                    "                  <div class=\"caption\">\n" +
                    "                  <p>\n" +
                    "                  <a  href=\"/article/"+ value.Slug +"\" class=\"vb-tooltip\" style=\"text-decoration:none; color: black;font-size: 16px\" data-toggle=\"tooltip\" title=\""+ value.Title +" +\">" + value.Title + "</a>\n" +
                    "                  </p>\n" +
                    "                  </div>\n" +
                    "                  </div>\n" +
                    "                  </div>\n"
          });

          $(".vb-history.row").html(vb_history);
        }
      });
    });

    // 下一页
    $(".vb-next-page").click(function(){
      const next_default_page = $("[next-default-page]").attr("next-default-page");

      const next_page_class = $(".vb-next-page").attr("class");

      if(next_page_class.indexOf("disabled") >= 0 ) {
        return false;
      }

      let w_url = window.location.href.split("/category/");
      let category_id = w_url[1] ? w_url[1] : '';
      let filter = {
        'category_id' : category_id,
        'p' : next_default_page
      };

      $.ajax({
        url:"/pagination",
        data:filter,
        success:function(result){
          $(".vb-last-page.disabled").attr("class","vb-last-page");
          $(".vb-last-page").attr("last-default-page",Number(result.PageNo) - 1);
          if ( !result.LastPage )
          {
             $(".vb-next-page").attr("next-default-page",Number(result.PageNo) + 1);
          }else{
            // 最后一页禁用下一页
            $(".vb-next-page").attr("class","vb-next-page disabled");
          }

          let vb_history = '';
          $.each(result.List,function(index,value) {
             vb_history += "<div class=\"col-sm-6 col-md-3\" >\n" +
                     "                  <div class=\"thumbnail\">\n" +
                     "                  <a href=\"/article/" + value.Slug + "\">\n" +
                     "                  <div class=\"vb-thumb-img\"  style=\"width: 170px;\n" +
                     "          height: 100px;\n" +
                     "          min-width: 100%;\n" +
                     "          min-height: 100%;\n" +
                     "          background-repeat: no-repeat;\n" +
                     "          background-position: center center;\n" +
                     "          background-size: cover;\n" +
                     "          background-image:url("+ value.VideoCover + ")\">\n" +
                     "                  </div>\n" +
                     "                  </a>\n" +
                     "                  <div class=\"caption\">\n" +
                     "                  <p>\n" +
                     "                  <a  href=\"/article/"+ value.Slug +"\" class=\"vb-tooltip\" style=\"text-decoration:none; color: black;font-size: 16px\" data-toggle=\"tooltip\" title=\""+ value.Title +" +\">" + value.Title + "</a>\n" +
                     "                  </p>\n" +
                     "                  </div>\n" +
                     "                  </div>\n" +
                     "                  </div>\n"
          });

          $(".vb-history.row").html(vb_history);

        }
      });
    });
  });

</script>

</body>
</html>
