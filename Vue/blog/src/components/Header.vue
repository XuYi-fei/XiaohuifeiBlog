<template>
  <header class="home headers">
    <div id="header-parent" style="flex-wrap: nowrap">
      <a href="/" id="header-title-a" style="white-space: nowrap"> <span id="header-title"> 小灰飞's小破站 </span> </a>
      <ul id="header-ul">
        <li><a href="/home" class="header-a">首页</a></li>
        <li><a href="/classification" class="header-a">分类</a></li>
        <li><a href="/comments" class="header-a">评论</a></li>
        <li><a href="/sentences" class="header-a">美句</a></li>
        <li><a href="#" class="header-a">关于我</a></li>
      </ul>
      <div style="display: inline-block" id="login-register">
        <div id="login-div" style="">
          <div class="function-buttons" v-show="!login"><LoginDialog> </LoginDialog></div>
          <div class="function-buttons" v-show="!login"><RegisterDialog> </RegisterDialog></div>
          <el-button type="text" @click="personal_space()" show-close="true" v-show="login"> 个人空间 </el-button>
          <el-button type="text" @click="enterAdmin()" show-close="true" v-show="admin"> 进入后台 </el-button>
          <el-button type="text" @click="cancel()" show-close="true" v-show="login"> 注销 </el-button>
        </div>
      </div>

    </div>
    <div id="divider">
      <el-divider></el-divider>
    </div>

  </header>


</template>

<script>
import LoginDialog from "./login/LoginDialog";
import RegisterDialog from "./register/RegisterDialog";
import {loginCheck} from "../api/login";
export default {
  name: "Header",
  data(){
    return {
      login: false,
      userId: '',
      admin: false
    }
  },
  methods:{
    enterAdmin(){
      window.location = '/admin'
    },
    cancel(){
      window.localStorage.removeItem('token')
      window.location.reload()
    },
    personal_space(){
      window.location = '/personal/'+ this.userId + '/personal-info'
    }
  },
  async created() {
    const response = await loginCheck()
    try{
        if(response.status === 0){
          this.login = true
            this.userId = response.data.userId
        }
        if(response.data.privilege === 1){
          this.admin = true
        }
      }catch (e) {}
  },
  watch: {
    $route: {
      async handler(val, oldVal){
        const response = await loginCheck()
        try{
            if(response.status === 0)
              this.login = true
            if(response.privilege === 1){
              this.admin = true
            }
          }catch (e) {
          }
        }
      }
  },
  components: {RegisterDialog, LoginDialog},
}
</script>

<style scoped>
/* 全局设置 */
a{
  text-decoration: none;
}



/* 组件设置 */

/* 首页导航栏 */
div#header-parent{
  display: flex;
  justify-content: space-around;
  background-color: black;
  min-width: 1440px;
}


/* 首页标题部分 */
a#header-title-a{
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex: 15%
}

span#header-title{
  font-family: "Microsoft YaHei UI",serif;
  color: lightskyblue;
  font-size: xx-large;
}

/* 首页链接部分 */
ul#header-ul{
  list-style: none;
  display: inline-block;
  min-width: 500px;
  white-space: nowrap;
  position: relative;
  top: 10px;
  flex: 60%
}

li{
  display: inline;
  margin: 20px;
}

a.header-a{
  color: greenyellow;
  font-size: large;
}

/* 登录与注册 */
div#login-register{
  display: inline;
  flex: 25%;
  min-width: 80px;
  white-space: nowrap;

}
div#login-div{
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}

div.function-buttons{
  display: flex;
  flex-direction: column;
  justify-content: center;
  margin: 5px 1px 5px 1px;
}

</style>
