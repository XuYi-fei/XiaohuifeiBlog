<template>
  <vue-particles
      color="#fff"
      :particleOpacity="0.7"
      :particlesNumber="60"
      shapeType="circle"
      :particleSize="4"
      linesColor="#fff"
      :linesWidth="2"
      :lineLinked="true"
      :lineOpacity="0.4"
      :linesDistance="150"
      :moveSpeed="2"
      :hoverEffect="true"
      hoverMode="grab"
      :clickEffect="true"
      clickMode="push"
      class="lizi"
      style="height: 1600px;min-width: 1440px; background: rgb(20 ,20 ,20); width: 100vw;"
  >
  </vue-particles>
  <div style="min-width: 1440px; width: 100vw; position: absolute; top: 0">
    <Header></Header>
    <Footer></Footer>
  </div>

</template>

<script>
import Header from "../components/Header";
import Footer from "../components/Footer";
export default {
  name: "Comments",
  components: {Footer, Header}
}
</script>

<style scoped>

</style>
