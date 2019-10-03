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
            {{ scope.row.id }}
          </template>
        </el-table-column>

        <el-table-column label="类别名字">
          <template slot-scope="scope">
            {{ scope.row.name }}
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
    <el-dialog title="添加文章" :visible.sync="dialogCreateFormVisible"  class="edit-rule-form">
      <el-form :model="createForm" :rules="rules" ref="createForm">

        <el-form-item label="类别名字" :label-width="formLabelWidth" prop="name">
          <el-input v-model="createForm.name" autocomplete="off"></el-input>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogCreateFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitCreateForm('createForm')">确 定</el-button>
      </div>
    </el-dialog>

<!--  编辑表单  -->
    <el-dialog title="修改类别" :visible.sync="dialogFormVisible" class="edit-rule-form">
      <el-form :model="editForm" :rules="rules" ref="editForm">

        <el-form-item label="类别名字" :label-width="formLabelWidth" prop="name">
          <el-input v-model="editForm.name" autocomplete="off"></el-input>
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
  import { index,store,update,destroy } from '@/api/categories.js'

  export default {

    methods: {

      handleSizeChange(val) {
        this.listLoading = true
        index({
          page : this.currentPage,
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
        this.editForm.category_id = row.id
        this.editForm.name = row.name
      },

      handleDelete(index, row) {
        this.$confirm('此操作将删除该分类, 是否继续?', '提示', {
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

      // 提交新文章 submitCreateForm
      submitCreateForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            let createFormData = {
              'name' : this.createForm.name,
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
              'id' : this.editForm.category_id,
              'name' : this.editForm.name,
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

        categories : [],
        tags : [],

        editListLoading: true,
        dialogFormVisible: false,
        dialogCreateFormVisible:false,

        editForm: {
          category_id: 0,
          name: "",
        },

        createForm: {
          category_id: 0,
          name: "",
        },

        formLabelWidth: '120px',

        rules: {
          name: [
            { required: true, message: '请输入名字', trigger: 'blur' },
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
