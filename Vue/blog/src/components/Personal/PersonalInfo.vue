<template>
  <div id="main">
    <div id="name-avatar-div">
      <div class="block"><el-avatar :size="60" :src="circleUrl"></el-avatar></div>
      <div class="name">
        <span class="name-id"> <span class="name-id-tips">账号Id:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; </span> {{ userInfo.userId }} </span>
        <span class="name-id"> <span class="name-id-tips">用户昵称:&nbsp;</span> {{ userInfo.userName }} </span>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive, toRefs } from 'vue';
import {getBlogByUser} from "../../api/editBlog";
import {loginCheck} from "../../api/login";

export default {
  name: "PersonalInfo",
  data(){
    return{
      userInfo: {}
    }
  },
  setup() {
    const state = reactive({
      circleUrl: 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
    });
    return {
      ...toRefs(state),
    };
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
  width: 700px;
  overflow: auto;
  display: flex;
  flex-direction: column;

}

div#name-avatar-div{
  display: flex;
  height: 200px;
}

div.block{
  padding: 10px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

div.name{
  display: flex;
  flex-direction: column;
  justify-content: center;
  text-align: left;
}

span.name-id{
  display: block;
  font-size: large;
}

span.name-id-tips{
  font-size: large;
  color: #3485ff;
}

</style>
