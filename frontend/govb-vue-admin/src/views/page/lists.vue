<template>
  <div>

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
            {{ scope.row.id }}
          </template>
        </el-table-column>

        <el-table-column label="名称">
          <template slot-scope="scope">
            {{ scope.row.display_name }}
          </template>
        </el-table-column>

        <el-table-column label="url" >
          <template slot-scope="scope">
            {{ scope.row.name }}
          </template>
        </el-table-column>

        <el-table-column label="排序" >
          <template slot-scope="scope">
            {{ scope.row.sort_order }}
          </template>
        </el-table-column>


        <el-table-column align="center" prop="created_at" label="添加时间" width="200">
          <template slot-scope="scope">
            <i class="el-icon-time"/>
            <span>{{ scope.row.created_at }}</span>
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
    <el-dialog title="添加页面" :visible.sync="dialogCreateFormVisible"  class="edit-rule-form">
      <el-form :model="createForm" :rules="rules" ref="createForm">

        <el-form-item label="页面名称" :label-width="formLabelWidth" prop="display_name">
          <el-input v-model="createForm.display_name" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="页面uri" :label-width="formLabelWidth" prop="name">
          <el-input v-model="createForm.name" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="页面内容" :label-width="formLabelWidth" prop="content">
          <mavon-editor style="height: 100%" v-model="createForm.content" ></mavon-editor>
        </el-form-item>

        <el-form-item label="显示顺序"  :label-width="formLabelWidth">
          <el-input-number
            v-model="createForm.sort_order"
            label="显示顺序"></el-input-number>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogCreateFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitCreateForm('createForm')">确 定</el-button>
      </div>
    </el-dialog>

<!--  编辑表单  -->
    <el-dialog title="修改页面" :visible.sync="dialogFormVisible" class="edit-rule-form">
      <el-form :model="editForm" :rules="rules" ref="editForm">

        <el-form-item label="页面名称" :label-width="formLabelWidth" prop="display_name">
          <el-input v-model="editForm.display_name" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="页面uri" :label-width="formLabelWidth" prop="name">
          <el-input v-model="editForm.name" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="页面内容" :label-width="formLabelWidth" prop="content">
          <mavon-editor style="height: 100%" v-model="editForm.content" ></mavon-editor>
        </el-form-item>

        <el-form-item label="显示顺序"  :label-width="formLabelWidth">
          <el-input-number
            v-model="editForm.sort_order"
            label="显示顺序"></el-input-number>
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
  import { index,update,destroy,store } from '@/api/pages'
  import { mavonEditor } from 'mavon-editor'
  import 'mavon-editor/dist/css/index.css'

  export default {

    components:{
      mavonEditor
    },

    methods: {

      handleSizeChange(val) {
        this.listLoading = true
        index({
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
        index({
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
        index({
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
        this.editForm.page_id = response_obj.id
        this.editForm.name = response_obj.name
        this.editForm.display_name = response_obj.display_name
        this.editForm.content = response_obj.content
        this.editForm.sort_order = parseInt(response_obj.sort_order)
      },

      handleDelete(index, row) {
          this.$confirm('此操作将删除该页面, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
          }).then(() => {
              destroy({
                  id : row.id
              }).then(response => {
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

      fetchData() {
        this.listLoading = true
        index(
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
      },

      // 提交新页面 submitCreateForm
      submitCreateForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            let createFormData = {
              'name' : this.createForm.name,
              'display_name' : this.createForm.display_name,
              'content' : this.createForm.content,
              'sort_order': this.createForm.sort_order,
            }

            store(createFormData).then(response => {
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
          if (valid) {
            let updateFormData = {
              'id' : this.editForm.page_id,
              'name' : this.editForm.name,
              'display_name' : this.editForm.display_name,
              'content' : this.editForm.content,
              'sort_order': this.editForm.sort_order,
            }

            update(updateFormData).then(response => {
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

        editListLoading: true,
        dialogFormVisible: false,
        dialogCreateFormVisible:false,

        editForm: {
          page_id: 0,
          name: '',
          display_name: '',
          content: '',
          sort_order: 0,

        },

        createForm: {
          name: '',
          display_name: '',
          content: '',
          sort_order: 0,
        },

        formLabelWidth: '120px',

        rules: {
          name: [
            { required: true, message: '请输标题', trigger: 'blur' },
          ],
          display_name: [
            { required: true, message: '请输描述', trigger: 'blur' },
          ],
          content: [
            { required: true, message: '请输入内容', trigger: 'blur' },
          ],
        }
      };
    },
    created() {
      this.fetchData()
    },

  }
</script>
