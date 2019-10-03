<template>
  <div>

    <div class="blog-search-div">
      搜索： <span class="search-input">
      <el-input
        placeholder="请输入内容"
        prefix-icon="el-icon-search"
        @blur="handleSearch"
        size="small"
        v-model="keyword">
      </el-input>
    </span>
    </div>

    <div style="
    padding-right: 20px;
    padding-top: 10px;float: right; display: inline" >
    <el-button
      type="primary"
      icon="el-icon-plus"
      circle
      @click="handleCreate"
    ></el-button>
    </div>

    <div class="app-container">
      <el-table
        v-loading="listLoading"
        :data="list"
        element-loading-text="Loading"
        border
        fit
        highlight-current-row>

        <el-table-column align="center" label="ID" width="95">
          <template slot-scope="scope">
            {{ scope.row.Id }}
          </template>
        </el-table-column>

        <el-table-column label="类别">
          <template slot-scope="scope">
            {{ scope.row.CategoryName }}
          </template>
        </el-table-column>

        <el-table-column label="标题" >
          <template slot-scope="scope">
            {{ scope.row.Title }}
          </template>
        </el-table-column>

        <el-table-column label="封面图" >
          <template slot-scope="scope">
            <img :src=scope.row.VideoCover width="40px" height="40px">
          </template>
        </el-table-column>

        <el-table-column align="center" prop="created_at" label="添加时间" width="200">
          <template slot-scope="scope">
            <i class="el-icon-time"/>
            <span>{{ scope.row.CreatedAt }}</span>
          </template>
        </el-table-column>

        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              size="mini"
              @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>

      </el-table>

      <p></p>

      <el-pagination
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="[10, 20, 30, 40]"
        :page-size="limit"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
      </el-pagination>

    </div>

<!--    添加表单-->
    <el-dialog title="添加文章" :visible.sync="dialogCreateFormVisible"  class="edit-rule-form">
      <el-form :model="createForm" :rules="rules" ref="createForm">

        <el-form-item label="标题" :label-width="formLabelWidth" prop="title">
          <el-input v-model="createForm.title" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="自定义视频封面图链接">
          <el-switch v-model="customize_preview_img" @change="switchChange"></el-switch>
        </el-form-item>

        <el-form-item label="视频封面图"  :label-width="formLabelWidth" v-if="createFormVisiblePreviewImg" >
           <el-input v-model="createForm.preview_img" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="视频封面图" :label-width="formLabelWidth"  v-if="createFormVisibleElUpload" >
          <el-upload
            class="upload-demo"
            action="/sm.ms/api/upload"
            name="smfile"
            :on-preview="handleCreatePreviewImg"
            :on-success="handleCreateUploadPreviewImg"
            :on-remove="handleCreateRemove"
            :file-list="createForm.preview_img_show"
            :before-upload="beforeCreatePreviewImgUpload"
            list-type="picture">
            <el-button size="small" type="primary">点击上传</el-button>
            <div slot="tip" class="el-upload__tip">不超过5M</div>
          </el-upload>
        </el-form-item>

        <el-form-item label="文章slug" :label-width="formLabelWidth" prop="slug">
          <el-input v-model="createForm.slug" autocomplete="off" placeholder="留空自动生成"></el-input>
        </el-form-item>

        <el-form-item label="类型" :label-width="formLabelWidth" prop="category_id">
          <el-select v-model="createForm.category_id" placeholder="请选择文章类型">
            <el-option
              v-for="item in this.categories"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否推荐" :label-width="formLabelWidth" prop="is_recommend">
          <el-radio v-model="createForm.is_recommend" label="0">否</el-radio>
          <el-radio v-model="createForm.is_recommend" label="1">是</el-radio>
        </el-form-item>


        <el-form-item label="视频地址" :label-width="formLabelWidth">
          <el-input v-model="createForm.video" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="文章内容" :label-width="formLabelWidth" prop="content">
          <mavon-editor style="height: 100%" v-model="createForm.content" ref=md @imgAdd="$imgAdd" @imgDel="$imgDel"></mavon-editor>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogCreateFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitCreateForm('createForm')">确 定</el-button>
      </div>
    </el-dialog>

<!--  编辑表单  -->
    <el-dialog title="修改文章" :visible.sync="dialogFormVisible" class="edit-rule-form">
      <el-form :model="editForm" :rules="rules" ref="editForm">

        <el-form-item label="标题" :label-width="formLabelWidth" prop="title">
          <el-input v-model="editForm.title" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="自定义视频封面图链接">
          <el-switch v-model="customize_preview_img" @change="switchChange"></el-switch>
        </el-form-item>

        <el-form-item label="视频封面图"  :label-width="formLabelWidth" v-if="editFormVisiblePreviewImg" >
          <el-input v-model="editForm.preview_img" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="视频封面图" :label-width="formLabelWidth" v-if="editFormVisibleElUpload">
          <el-upload
            class="upload-demo"
            action="/sm.ms/api/upload"
            name="smfile"
            :on-preview="handlePreviewImg"
            :on-success="handleUploadPreviewImg"
            :on-remove="handleRemove"
            :file-list="editForm.preview_img_show"
            :before-upload="beforePreviewImgUpload"
            list-type="picture">
            <el-button size="small" type="primary">点击上传</el-button>
            <div slot="tip" class="el-upload__tip">不超过5M</div>
          </el-upload>
        </el-form-item>

        <el-form-item label="文章slug" :label-width="formLabelWidth" prop="slug">
          <el-input v-model="editForm.slug" autocomplete="off" placeholder="留空自动生成"></el-input>
        </el-form-item>

        <el-form-item label="类型" :label-width="formLabelWidth" prop="category_id">
          <el-select v-model="editForm.category_id" placeholder="请选择文章类型">
            <el-option
              v-for="item in this.categories"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否推荐" :label-width="formLabelWidth" prop="is_recommend">
          <el-radio v-model="editForm.is_recommend" label="0">否</el-radio>
          <el-radio v-model="editForm.is_recommend" label="1">是</el-radio>
        </el-form-item>

        <el-form-item label="视频地址" :label-width="formLabelWidth">
          <el-input v-model="editForm.video" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="文章内容" :label-width="formLabelWidth" prop="content">
          <mavon-editor
            style="height: 100%"
            v-model="editForm.content"
            ref=md @imgAdd="$imgAdd" @imgDel="$imgDel"
          ></mavon-editor>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitEditForm('editForm')">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
  import { getPosts,delPosts,storePosts,updatePosts } from '@/api/posts'
  import { getCategories } from '@/api/categories'
  import { mavonEditor } from 'mavon-editor'
  import 'mavon-editor/dist/css/index.css'
  import { mavonUpload } from '@/api/common'

  export default {

    components:{
      mavonEditor
    },

    filters: {
      statusFilter(status) {
        const statusMap = {
          published: 'success',
          draft: 'gray',
          deleted: 'danger'
        }
        return statusMap[status]
      }
    },

    methods: {

      switchChange(status){
        this.createFormVisiblePreviewImg = !this.createFormVisiblePreviewImg
        this.createFormVisibleElUpload = !this.createFormVisibleElUpload
        this.editFormVisibleElUpload = !this.editFormVisibleElUpload
        this.editFormVisiblePreviewImg = !this.editFormVisiblePreviewImg
      },

      $imgAdd(pos, $file) {
        // 第一步.将图片上传到服务器.
        let formdata = new FormData();
        formdata.append('smfile', $file);

        mavonUpload(formdata).then((res) => {
          if ( res.code == 'success' )
          {
            this.$refs.md.$img2Url(pos, res.data.url);
          }
        }).catch(e => {
          alert(e);
        });
      },

      $imgDel(pos) {
        delete this.img_file[pos];
      },

      handleSizeChange(val) {
        this.listLoading = true
        getPosts({
          limit : val,
          page : this.currentPage,
          keyword : this.keyword,
        }).then(response => {
          this.list = response.data.List
          this.listLoading = false
          this.currentPage = response.data.PageNo
          this.total = response.data.TotalCount
          this.limit = parseInt(response.data.PageSize)
        })
      },

      handleCurrentChange(val) {
        this.listLoading = true
        getPosts({
          page : val,
          limit : this.limit,
          keyword : this.keyword,
        }).then(response => {
          this.list = response.data.List
          this.listLoading = false
          this.currentPage = response.data.PageNo
          this.total = response.data.TotalCount
          this.limit = parseInt(response.data.PageSize)
        })
      },

      handleSearch() {
        this.listLoading = true
        getPosts({
          keyword : this.keyword,
          limit : this.limit
        }).then(response => {
          this.list = response.data.List
          this.listLoading = false
          this.currentPage = response.data.PageNo
          this.total = response.data.TotalCount
          this.limit = parseInt(response.data.PageSize)
        })
      },

      handleCreate(){
        this.dialogCreateFormVisible = true
      },

      handleEdit(index, row) {
        this.dialogFormVisible = true
        let response_obj = row;
        this.editForm.post_id = response_obj.Id
        this.editForm.category_id = response_obj.CategoryId
        this.editForm.title = response_obj.Title
        this.editForm.video = response_obj.VideoUrl
        this.editForm.is_recommend = response_obj.IsRecommend
        this.editForm.preview_img = response_obj.VideoCover
        this.editForm.preview_img_show = [
          {
            'name' : response_obj.VideoCover,
            'url' : response_obj.VideoCover
          }
        ];
        this.editForm.html_content = response_obj.HtmlContent
        this.editForm.content = response_obj.Content
        this.editForm.slug = response_obj.Slug
      },

      handleDelete(index, row) {
        this.$confirm('此操作将删除该文章, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.listLoading = true
          delPosts({
            id : row.Id
          }).then(response => {
            this.listLoading = false
            this.$message({
              type: 'success',
              message: '删除成功!'
            });
            this.fetchData();
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });
        });
      },

      handleRemove(file, fileList) {
        this.editForm.preview_img_show = [];
        console.log(file, fileList);
      },

      handlePreviewImg(file) {
        console.log(file);
      },

      handleUploadPreviewImg(res, file){
        if ( res.code == "success" )
        {
          this.editForm.preview_img = res.data.url;
          this.editForm.preview_img_show = [
            {
              'name' : res.data.storename,
              'url' : res.data.url
            }
          ];
        }else{
          this.$message.error(res.message);
          this.editForm.preview_img_show = [];
          this.editForm.preview_img = '';
        }
      },

      beforePreviewImgUpload(file) {

        // 判断个数
        if ( this.editForm.preview_img_show.length >= 1 )
        {
          this.$message.error('只能上传一张');
        }

        const isJPG = file.type === 'image/jpeg';
        const isGIF = file.type === 'image/gif';
        const isPNG = file.type === 'image/png';
        const isBMP = file.type === 'image/bmp';

        if (!isJPG && !isGIF && !isPNG && !isBMP) {
          this.$message.error('上传图片必须是JPG/GIF/PNG/BMP 格式!');
        }

        const isLt5M = file.size / 1024 / 1024 / 1024 / 1024 / 1024 < 5;

        if (!isLt5M) {
          this.$message.error('上传头像图片大小不能超过 5MB!');
        }
        return isLt5M;
      },

      handleCreateRemove(file, fileList) {
        this.editForm.preview_img_show = [];
        console.log(file, fileList);
      },

      handleCreatePreviewImg(file) {
        console.log("asdasdada");
        console.log(file);
      },

      handleCreateUploadPreviewImg(res, file){
        if ( res.code == "success" )
        {
          this.createForm.preview_img = res.data.url;
          this.createForm.preview_img_show = [
            {
              'name' : res.data.storename,
              'url' : res.data.url
            }
          ];
        }else{
          // sm.ms 更新了 同一张图片会限制上传次数
          this.$message.error(res.message);
          this.createForm.preview_img_show = [];
          this.createForm.preview_img = ''
        }
      },

      beforeCreatePreviewImgUpload(file) {

        // 判断个数
        if ( this.createForm.preview_img_show.length >= 1 )
        {
          this.$message.error('只能上传一张');
        }

        const isJPG = file.type === 'image/jpeg';
        const isGIF = file.type === 'image/gif';
        const isPNG = file.type === 'image/png';
        const isBMP = file.type === 'image/bmp';

        if (!isJPG && !isGIF && !isPNG && !isBMP) {
          this.$message.error('上传图片必须是JPG/GIF/PNG/BMP 格式!');
        }

        const isLt5M = file.size / 1024 / 1024 / 1024 / 1024 / 1024 < 5;

        if (!isLt5M) {
          this.$message.error('上传头像图片大小不能超过 5MB!');
        }
        return isLt5M;
      },

      fetchData() {
        this.listLoading = true
        getPosts(
          {
            page : this.currentPage,
            limit : this.limit,
            keyword : this.keyword,
          }
        ).then(response => {

          this.list = response.data.List
          this.listLoading = false
          this.currentPage = response.data.PageNo
          this.total = response.data.TotalCount
          this.limit = parseInt(response.data.PageSize)

        })

        getCategories().then(response => {
          this.categories = response.data
          this.createForm.category_id = response.data[0].id
        })
      },

      // 提交新文章 submitCreateForm
      submitCreateForm(formName) {
        this.$refs[formName].validate((valid) => {
          if ( this.createForm.preview_img === "" )
          {
            this.$message.error('封面图必须上传');
            return false;
          }
          if (valid) {
            let createFormData = {
              'title' : this.createForm.title,
              'slug' : this.createForm.slug,
              'content' : this.createForm.content,
              'video_url' : this.createForm.video,
              'video_cover' : this.createForm.preview_img,
              'category_id' : this.createForm.category_id,
              'is_recommend' : this.createForm.is_recommend,
            }

            storePosts(createFormData).then(response => {
              this.dialogCreateFormVisible = false;
              this.$message({
                type: 'success',
                message: '添加成功!'
              });
              this.fetchData();
              this.$refs[formName].resetFields(); // 重置表单
            })
          } else {
            this.dialogCreateFormVisible = true;
            return false;
          }
        });
      },

      // 提交修改表单
      submitEditForm(formName) {
        this.$refs[formName].validate((valid) => {
          if ( this.editForm.preview_img === "" )
          {
            this.$message.error('封面图必须上传');
            return false;
          }
          if (valid) {
            let updateFormData = {
              'id' : this.editForm.post_id,
              'title' : this.editForm.title,
              'slug' : this.editForm.slug,
              'content' : this.editForm.content,
              'video_url' : this.editForm.video,
              'video_cover' : this.editForm.preview_img,
              'category_id' : this.editForm.category_id,
              'is_recommend' : this.editForm.is_recommend,
            }

            updatePosts(updateFormData).then(response => {
              this.dialogFormVisible = false;
              this.$message({
                type: 'success',
                message: '修改成功!'
              });
              this.fetchData();
              this.$refs[formName].resetFields(); // 重置表单
            })
          } else {
            this.dialogFormVisible = true;
            return false;
          }
        });
      },

    },

    data() {
      return {
        currentPage: 1,
        total: 0,
        limit: 10,
        keyword : '',
        list: null,
        listLoading: true,

        categories : [],
        tags : [],
        img_file : [],

        editListLoading: true,
        dialogFormVisible: false,
        dialogCreateFormVisible:false,

        createFormVisibleElUpload: true,
        editFormVisibleElUpload: true,

        createFormVisiblePreviewImg: false,
        editFormVisiblePreviewImg: false,

        smUploadHeader:{
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },

        editForm: {
          post_id: 0,
          title: '',
          description: '',
          slug: '',
          content: '',
          html_content: '',
          video: '',
          preview_img: '',
          preview_img_show: [],
          category_id: 0,
          is_recommend: "0",
        },

        customize_preview_img: false,

        createForm: {
          title: '',
          description: '',
          slug: '',
          content: '',
          html_content: '',
          video: '',
          preview_img: '',
          preview_img_show: [],
          category_id: 0,
          is_recommend: "0",
        },

        formLabelWidth: '120px',

        rules: {
          title: [
            { required: true, message: '请输标题', trigger: 'blur' },
          ],
          description: [
            { required: true, message: '请输描述', trigger: 'blur' },
          ],
          // slug: [
          //   { required: true, message: '请输入slug', trigger: 'blur' },
          // ],
          content: [
            { required: true, message: '请输入内容', trigger: 'blur' },
          ],
          category_id: [
            { required: true, message: '请选择活动区域', trigger: 'change' }
          ],
          video: [
            { required: true, message: '视频地址必填', trigger: 'blur' }
          ],
        }
      };
    },
    created() {
      this.fetchData()
    },

  }
</script>

<style>
  .blog-search-div {
    padding-left: 20px;
    padding-right: 20px;
    padding-top: 10px;
    margin: 0 0 20px 0;
    float: left;
  }
  .search-input {
    display: inline-flex;
    width: 300px;
  }
</style>
