<template>
  <div id="main">
    <div class="information-div">
      <div class="information-title">
        <span> 基本信息 </span>
      </div>
      <el-divider style="margin: 1px 0 1px 0"></el-divider>
      <div id="base-info-content">
        <div id="base-info-content-show"></div>
        <ul>
          <li> <span class="info-key">用户 &nbsp;&nbsp; ID</span> <span class="info-value">{{ userInfo.userId }}</span></li>
          <li> <span class="info-key">用户昵称</span> <span class="info-value">{{ userInfo.userName }}</span></li>
          <li> <span class="info-key">性&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;别</span> <span class="info-value">{{ userInfo.sex }}</span></li>
          <li> <span class="info-key">年&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;龄</span> <span class="info-value">{{ userInfo.age }}</span></li>
          <li> <span class="info-key">真实姓名</span> <span class="info-value">{{ userInfo.realName }}</span></li>
        </ul>
      </div>
    </div>

    <div class="information-div">
      <div class="information-title">
        <span> 基本信息 </span>
      </div>
      <el-divider style="margin: 1px 0 1px 0"></el-divider>
      <div id="blog-info-content">
        <div id="blog-info-content-show"></div>
        <ul>
          <li> <span class="info-key">博客数量</span> <span class="info-value">{{ userInfo.userId }}</span></li>
          <li> <span class="info-key">用户昵称</span> <span class="info-value">{{ userInfo.userName }}</span></li>
          <li> <span class="info-key">性&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;别</span> <span class="info-value">{{ userInfo.sex }}</span></li>
          <li> <span class="info-key">年&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;龄</span> <span class="info-value">{{ userInfo.age }}</span></li>
          <li> <span class="info-key">真实姓名</span> <span class="info-value">{{ userInfo.realName }}</span></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import {loginCheck} from "../../api/login";

export default {
  name: "PersonalDetail",
  data(){
    return{
      userInfo: {}
    }
  },
  async created() {
    const response = await loginCheck()
    if (response.status !== 0){
      alert("登录失效，请重新登录")
      window.location = '/home'
    }else{
      this.userInfo = response.data
    }
  }
}
</script>

<style scoped>
div#main{
  display: flex;
  flex-direction: column;
  text-align: left;
}
div.information-div{
  display: flex;
  flex-direction: column;
  background-color: white;
}

div.information-div:hover{
  background-color: #f8f7ff;
}

div.information-title{
  height: 48px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: black;
  font-weight: bold;
  font-size: 20px;
  padding-left: 10px;
}

div#base-info-content{
  padding: 16px;
}

div#base-info-content-show{
  padding: 8px 0 8px 0;
}

ul{
  list-style: none;
}

li{
  margin-bottom: 28px ;
}

li span.info-key{
  font-size: large;
  padding-right: 20px;
  color: #555666;
}

li span.info-value{
  font-size: large;
  padding-left: 20px;
}
</style>
