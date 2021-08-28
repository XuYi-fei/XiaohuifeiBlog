<template>
  <div id="footer">
    <a href="https://icp.chinaz.com/xuyi-fei.xyz"> 鄂ICP备2021002829号 </a>
  </div>
</template>

<script>
export default {
  name: "Footer"
}
</script>

<style scoped>
a{
  color: #9dff7c;

}
div#footer{
  position: fixed;
  right: 0;
  top: 50vm;
  background-color: #000;
  height: 50px;
  z-index: 30000;
  line-height: 40px;
  text-align: center;
  font-size: large;
  font-weight: bold;
}
</style>
