<div class="vb-detail ">

    <div class="well well-sm">
        <article class="markdown-body">
            <h1 class="text-center">
                <b>
                    {{.PostDetail.Title}}
                </b>
            </h1>
            {{ str2html .PostDetail.HtmlContent }}

            {{if .PostDetail.VideoUrl}}
            <div class="vb-video">

                <video id="video" class="video-js vjs-default-skin vjs-fluid vjs-big-play-centered " poster="{{.PostDetail.VideoCover}}" width="375" height="200" controls preload="none"
                       data-setup='{ "html5" : { "nativeTextTracks" : false } }'>
                    <source src="{{.PostDetail.VideoUrl}}" type="video/mp4">
                    <p class="vjs-no-js">  播放视频需要启用 JavaScript，推荐使用<a href="http://videojs.com/html5-video-support/" target="_blank">支持HTML5</a>的浏览器访问。</p>
                </video>

            </div>
            {{end}}
        </article>
    </div>


</div>

<div class="page-header text-center">
    <h2>评论
        <small>说点什么吧！</small>
    </h2>
</div>

<div id="vcomments"></div>
<script src="//cdn.bootcss.com/video.js/7.6.0/alt/video.core.min.js"></script>
<script src='//unpkg.com/valine/dist/Valine.min.js'></script>
{{/*https://ws1.sinaimg.cn/large/006c53jqgy1g7gg2fkxx6j303e03kgle.jpg*/}}
<script>
    new Valine({
        el: '#vcomments',
        appId: 'leancloud-appid',
        appKey: 'leancloud-appKey',
        notify: false,
        verify: false,
        avatar: 'monsterid',
        placeholder: '多点思考，理性评论！'
    });
</script>

