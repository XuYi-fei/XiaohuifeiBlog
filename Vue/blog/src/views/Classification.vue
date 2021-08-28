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
      style="height: 1700px;min-width: 1440px; background: rgb(20 ,20 ,20); width: 100vw; z-index: 20000; float: left"
  >
  </vue-particles>
    <div id="main-menu" style="min-width: 1440px; width: 100vw; position: absolute; top: 0">
      <Header></Header>
      <div id="flex">
        <div id="field-div">
          <suspense>
            <template #default>
              <ClassificationPage />
            </template>
            <template #fallback>
              加载中....
            </template>
          </suspense>
        </div>
        <div id="blogs-div">
          <suspense>
            <template #default>
              <ClassificationContentPage />
            </template>
            <template #fallback>
              加载....
            </template>
          </suspense>
        </div>
        <div id="info-div">
          <suspense>
            <template #default>
              <ClassificationInfoPage />
            </template>
            <template #fallback>
              加载中....
            </template>
          </suspense>
        </div>

      </div>
      <Footer></Footer>
    </div>

</template>

<script>
import Header from "../components/Header";
import ClassificationPage from "./ClassificationPage";
import ClassificationContentPage from "./ClassificationContentPage";
import ClassificationInfoPage from "./ClassificationInfoPage";
import Footer from "../components/Footer";

export default {
  name: "Classification",
  components: {Footer, ClassificationInfoPage, ClassificationContentPage, ClassificationPage, Header},

}
</script>

<style scoped>

span.blog-tag{
  display: block;
  font-size: 22px;
}

div#flex{
  display: flex;
}

div#field-div{
  flex: 15%;
  display: flex;
  flex-direction: row-reverse;
  min-width: 165px;
  margin: 10px;
}

div#blogs-div{
  flex:60%;
}

div#info-div{
  flex: 25%;
}
</style>
