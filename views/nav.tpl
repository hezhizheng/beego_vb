<nav class="navbar navbar-inverse navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <div class="logo" onclick="window.open('/','_self')"></div>
        </div>
        <div id="navbar" class="collapse navbar-collapse">
            <ul class="nav navbar-nav">
{{/*                <li class="active"><a href="/">首页</a></li>*/}}

                {{range $i, $v := .PageAll}}

                    <li><a href="/{{$v.Name}}">{{ $v.DisplayName }}</a></li>
                {{end}}


            </ul>
        </div><!--/.nav-collapse -->
    </div>
</nav>
