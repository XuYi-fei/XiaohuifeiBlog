<template>
  <Header></Header>
  <div id="main">
    <div id="all-blogs">

    </div>

    <div id="other-infos">

    </div>
  </div>
</template>

<script>
import Header from "../components/Header";
export default {
  name: "AllBlogs",
  components: {Header}
}
</script>

<style scoped>
div#main{
  display: flex;
}

div#all-blogs{
  display: flex;
  flex-direction: column;
  flex: 70%
}

div#other-infos{
  display: flex;
  flex-direction: column;
  flex: 70%;
}
</style>
