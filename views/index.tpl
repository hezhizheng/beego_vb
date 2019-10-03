
<div class="row">

  <div class="col-sm-9">
    <div class="vb-content" >

      <div class="starter-template">

        <div class="page-header">
          <h1>推荐
            <small>热门</small>
          </h1>
        </div>

        <div id="myCarousel" class="carousel slide" >
          <!-- 轮播（Carousel）指标 -->
          <ol class="carousel-indicators">
            {{range $i, $v := .recommend}}

            <li data-target="#myCarousel" data-slide-to="{{$i}}"
                    {{if eq $i 0}}
                      class="active"
                    {{end}}
            ></li>
            {{end}}
          </ol>
          <!-- 轮播（Carousel）项目 -->
          <div class="carousel-inner">
            {{range $i, $v := .recommend}}

            <div class="item {{if eq $i 0}}
              active
{{end}}"
                   style="width: 800px;
                  height: 400px;
                  min-width: 100%;
                  min-height: 100%;
  background-repeat: no-repeat;
  background-position: center center;
  background-size: cover;
  cursor: pointer;
  background-image:url({{ $v.VideoCover }}"
                 onclick="window.open('/article/{{ $v.Slug }}','_self')">
              <div class="carousel-caption">{{ $v.Title }}</div>
              </div>
            {{end}}

          </div>
          <!-- 轮播（Carousel）导航 -->
          <a class="left carousel-control" href="#myCarousel" role="button" data-slide="prev">
            <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span>
            <span class="sr-only">Previous</span>
          </a>
          <a class="right carousel-control" href="#myCarousel" role="button" data-slide="next">
            <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span>
            <span class="sr-only">Next</span>
          </a>
        </div>

        <div class="page-header">
          <h1>过往记录
          </h1>
        </div>

        {{ if ne .Posts.TotalCount 0 }}
        <div class="vb-history row">
          {{range $i, $v := .Posts.List}}
            <div class="col-sm-6 col-md-3" >
              <div class="thumbnail">
                <a href="/article/{{ $v.Slug }}">
                  <div class="vb-thumb-img"  style="width: 170px;
                  height: 100px;
                  min-width: 100%;
                  min-height: 100%;
  background-repeat: no-repeat;
  background-position: center center;
  background-size: cover;
  background-image:url({{ $v.VideoCover }})">
                  </div>
                </a>
                <div class="caption">
                  <p>
                    <a  href="/article/{{ $v.Slug }}" class="vb-tooltip" style="text-decoration:none; color: black;font-size: 16px" data-toggle="tooltip" title="{{ $v.Title }}">{{ $v.Title }}</a>
                  </p>
                </div>
              </div>
            </div>
          {{end}}
        </div>
        {{else}}
          暂无...
        {{end}}

        {{ if gt .Posts.TotalPage 1}}
        <div class="row">
          <ul class="pager">
            <li class="vb-last-page disabled" last-default-page="1">
              <a href="javascript:;">
                &larr; 上一页
              </a>
            </li>
            &nbsp;&nbsp;&nbsp;&nbsp;
            <li class="vb-next-page {{if .Posts.LastPage }}
              disabled
{{end}}" next-default-page="2">
              <a href="javascript:;">
                下一页 &rarr;
              </a>
            </li>
          </ul>
        </div>
        {{end}}
      </div>

    </div>
  </div>

  <div class="col-sm-3">
    <div class="vb-side text-center">

      <div class="panel panel-default">
        <div class="panel-heading">
          <b>分类</b>
        </div>
        <div class="panel-body">
          <ul class="list-group">
            {{range $i, $v := .categories}}
              <a href="/category/{{$v.id}}" style="text-decoration:none">
            <li class="list-group-item">
              <span class="badge">{{$v.product_count}}</span>
              {{$v.name}}
            </li>
              </a>
            {{end}}
          </ul>
        </div>
      </div>

    </div>

{{/*    <div class="vb-side text-center">*/}}

{{/*      <div class="panel panel-default">*/}}
{{/*        <div class="panel-heading">*/}}
{{/*          <b>标签</b>*/}}
{{/*        </div>*/}}
{{/*        <div class="panel-body ">*/}}
{{/*          <ul class="vb-list-group">*/}}
{{/*            <li class="vb-list-labs-item"><span class="label label-info">新asfasf</span></li>*/}}
{{/*            <li class="vb-list-labs-item"><span class="label label-info">新asfasf</span></li>*/}}
{{/*            <li class="vb-list-labs-item"><span class="label label-info">新asfasf</span></li>*/}}
{{/*            <li class="vb-list-labs-item">*/}}
{{/*              <span class="label label-info">新24*7</span>*/}}
{{/*            </li>*/}}
{{/*            <li class="vb-list-labs-item"><span class="label label-info">新24*7</span></li>*/}}
{{/*            <li class="vb-list-labs-item">*/}}
{{/*              <span class="label label-info">新77</span>*/}}

{{/*            </li>*/}}
{{/*          </ul>*/}}
{{/*        </div>*/}}
{{/*      </div>*/}}

{{/*    </div>*/}}

  </div>

</div>






