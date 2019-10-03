<div class="vb-detail ">

    <div class="well well-sm">
        <article class="markdown-body">
            {{ str2html .Page.HtmlContent }}
        </article>
    </div>


</div>

<div class="page-header text-center">
    <h2>评论
        <small>说点什么吧！</small>
    </h2>
</div>

<div id="vcomments"></div>

<script src='//unpkg.com/valine/dist/Valine.min.js'></script>
{{/*https://ws1.sinaimg.cn/large/006c53jqgy1g7gg2fkxx6j303e03kgle.jpg*/}}
<script>
    new Valine({
        el: '#vcomments',
        appId: 'leancloud-appid',
        appKey: 'leancloud-appkey',
        notify: false,
        verify: false,
        avatar: 'monsterid',
        placeholder: '多点思考，理性评论！'
    });
</script>

