<template>
  <header>
    <span id="back-to-home">
      <el-button type="primary" @click="goBack()">退出后台</el-button>
    </span>
    <span id="back-title"> 小灰飞的后台 </span>
    <el-divider style="margin: 5px"></el-divider>
  </header>
  <div style="min-height: 980px">
    <el-tabs type="border-card" v-model="activeTab" @tab-click="changeTab">
      <el-tab-pane label="首页" name="index" @click="changeTab('index')"><router-view name="indexView"></router-view></el-tab-pane>
      <el-tab-pane label="美句审核" name="sentence" @click="changeTab('sentences')"><router-view name="sentencesView"></router-view></el-tab-pane>
      <el-tab-pane label="博客审核" name="blog"><router-view name="blogsView"></router-view></el-tab-pane>
      <el-tab-pane label="美句管理" name="sentenceAll"><router-view name="sentencesAllView"></router-view></el-tab-pane>
      <el-tab-pane label="博客管理" name="blogAll"><router-view name="blogsAllView"></router-view></el-tab-pane>
      <el-tab-pane label="申诉管理" name="complaint"><router-view name="complaintsView"></router-view></el-tab-pane>
    </el-tabs>
  </div>

  <Footer></Footer>
</template>

<script>
import Footer from "../Footer";
export default {
  name: "AdminHeader",
  components: {Footer},
  data(){
    return{
      activeTab: ''
    }
  },
  methods:{
    changeTab(e){
      window.location = '/admin/' + e.props.name
    },
    goBack(){
      window.location = '/home'
    }
  },
  created() {
    const path = this.$route.fullPath.split('/')
    this.activeTab = path[path.length-1]
    if(path[path.length-2] === 'checkblog'){
      this.activeTab = 'blog'

    }
  }
}
</script>

<style scoped>
header{
  min-width: 1440px;
  height: 80px;
  line-height: 80px;
  text-align: center;
  font-size: xx-large;
  color: #3485ff;
  font-family: "STXingkai", sans-serif;
}

span#back-to-home{
  padding-left: 10px;
  float: left;
  z-index: 100;
}

span#back-title{
  position: relative;
  right: 50px;
}
</style>
