<template>
  <div class="dashboard-container">
<!--    <div class="dashboard-text">name:{{ name }}</div>-->
<!--    <div class="dashboard-text">roles:<span v-for="role in roles" :key="role">{{ role }}</span></div>-->

    <div style="display: flex; margin-top: 20px; height: 100px;">

        <div class="transition-box">
          <router-link to="page/page">
          <svg-icon icon-class="icon_网页" />  页面：{{this.page_count}}
          </router-link>
        </div>

      <div class="transition-box">
        <router-link to="posts/lists">
        <svg-icon icon-class="book" />  文章：{{this.post_count}}
        </router-link>
      </div>

      <div class="transition-box">
        <router-link to="posts/categories">
        <svg-icon icon-class="folder-open" />  分类：{{this.category_count}}
        </router-link>
      </div>

    </div>

  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import { index } from '@/api/dashboards.js'

  export default {
    name: 'Dashboard',
    computed: {
      ...mapGetters([
        'name',
        'roles'
      ])
    },

    methods : {
      getCount() {
        index().then(response => {
          this.post_count = response.data.post_count
          this.category_count = response.data.category_count
          this.page_count = response.data.page_count
        })
      },
    },

    data() {
      return {
        post_count : 0,
        category_count : 0,
        page_count : 0,
      }
    },

    created() {
      this.getCount()
    },
  }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .dashboard {
    &-container {
      margin: 30px;
    }
    &-text {
      font-size: 30px;
      line-height: 46px;
    }
  }

  .transition-box {
    margin-bottom: 10px;
    width: 200px;
    height: 100px;
    border-radius: 4px;
    background-color: #304156;
    text-align: center;
    color: #fff;
    padding: 40px 20px;
    box-sizing: border-box;
    margin-right: 20px;
    cursor: pointer;
  }
</style>
