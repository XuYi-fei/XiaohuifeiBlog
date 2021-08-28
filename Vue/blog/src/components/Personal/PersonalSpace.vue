<template>
  <div id="main">
    <header>
      <div id="header-parent" style="flex-wrap: nowrap">
          <div data-v-b8276df4="" id="BreadCrumb">
            <div class="el-breadcrumb" aria-label="Breadcrumb" role="navigation" data-v-b8276df4="">
            <router-view name="bread"></router-view>
            </div>
          </div>

        <a href="/" id="header-title-a" style="white-space: nowrap"> <span id="header-title"> 小灰飞的个人空间 </span> </a>
        <ul id="header-ul">
          <li><a href="/home" class="header-a">回到首页</a></li>
        </ul>
      </div>
      <div id="divider">
        <el-divider></el-divider>
      </div>

    </header>
    <div style="display: flex;">
      <router-view name="nav"></router-view>
      <div style="flex: 80%; overflow: auto; height: 800px;">
        <router-view name="assets"></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import {loginCheck} from "../../api/login";
import Footer from "../Footer";

export default {
  name: "PersonalSpace",
  components: {Footer},
  data(){
    return{
      userId: ''
    }
  },
  setup(){
    const handleOpen = (key, keyPath) => {
      console.log(key, keyPath);
    };
    const handleClose = (key, keyPath) => {
      console.log(key, keyPath);
    };
    return {
      handleOpen,
      handleClose,
    };
  },
}
</script>

<style scoped>
div#main{
  min-width: 900px;
}

div.el-col el-col-4{
  min-width: 150px;
}
/* 全局设置 */
div#main{
  background: transparent;
  display: flex;
  flex-direction: column;

}

a{
  text-decoration: none;
}

header{
  padding-top: 8px;
  background-color: white;
  position: relative;
}

/* 组件设置 */
div#BreadCrumb{
  display: flex;
  justify-content: center;
  flex-direction: column;
  padding: 10px;
  flex: 25%;
}


/* 首页导航栏 */
div#header-parent{
  display: flex;
  justify-content: space-around;
}


/* 首页标题部分 */
a#header-title-a{
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex: 50%
}

span#header-title{
  font-family: "Microsoft YaHei UI",serif;
  color: dodgerblue;
  font-size: xx-large;
}

/* 首页链接部分 */
ul#header-ul{
  margin-top: 20px;
  list-style: none;
  display: inline;
  min-width: 200px;
  white-space: nowrap;
  flex: 25%
}

li{
  display: inline;
  margin: 20px;
}

a.header-a{
  color: dimgrey;
  font-size: large;
}
</style>
