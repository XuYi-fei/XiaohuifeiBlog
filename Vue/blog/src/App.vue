<template>
  <div>
    <router-view name="home"/>
    <router-view name="personal"/>
    <router-view name="blog"/>
    <router-link to="/"></router-link>
  </div>

</template>
<script>
</script>
<style lang="scss">
//html {
//  height: 100%;
//  background-image: -webkit-radial-gradient(ellipse farthest-corner at center top, #000d4d 0%, #000105 100%);
//  background-image: radial-gradient(ellipse farthest-corner at center top, #000d4d 0%, #000105 100%);
//  cursor: arrow ;
//}
//
//body {
//  width: 100%;
//  margin: 0;
//  overflow: auto;
//}

body{
  background-color: rgb(245,246,247);
  margin: 0 0 10px 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  //background-color: #5381ff;
  color: #2c3e50;
}

header.headers{
  padding-top: 8px;
  background-color: black;
  position: relative;
  z-index: 30000;
  min-width: 1440px;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;
    text-decoration: none;
    position: relative;
    z-index: 30000;

    &.router-link-exact-active {
      color: #42b983;
      text-decoration: none;
    }
  }
}
</style>
