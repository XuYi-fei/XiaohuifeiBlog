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
      style="height: 1400px;min-width: 1440px; background: rgba(20 ,20 ,20, 0); width: 100vw; position: absolute; top: 0; z-index: 20000"
  >
  </vue-particles>
  <div style="min-width: 1440px; width: 100vw;position: absolute; top: 0;">
    <Header></Header>
    <div id="main" style="background-color: black">
      <div id="carousel">
        <el-carousel :interval="3000" arrow="always" height="60px">
          <el-carousel-item v-for="item in sentenceToShowOri" :key="item">
            <div style="color: black"> {{ item.content}}
            </div>
          </el-carousel-item>
        </el-carousel>
      </div>
      <el-divider style="margin: 5px; background-color: #3485ff"></el-divider>
      <suspense>
        <template #default>
          <SentencePage />
        </template>
        <template #fallback>
          加载中...
        </template>
      </suspense>
    </div>
  </div>


</template>

<script>
import Header from "../components/Header";
import {loginCheck} from "../api/login";
import {defineAsyncComponent, reactive, ref, toRefs} from "vue";
import {getAllSentences, getSentencesPageNum, getSentencesToShow, submitSentence} from "../api/sentences";
import {UserFilled, Timer} from '@element-plus/icons'
import SentencePage from "../components/sentences/SentencePage";
export default {
  name: "Sentences",
  components: {
    SentencePage,
    Header,
  },
  async created(){
    const shows = await getSentencesToShow()
    if (shows.status === 0){
      this.sentencesToShow = shows.data
      this.sentenceToShowOri = this.sentenceToShowOri.concat(this.sentencesToShow)
    }
  },
  data() {
    return {
      sentenceToShowOri: [{content: '一花一世界，一叶一菩提'},],
      sentencesToShow: [],
    };
  },
}
</script>

<style scoped>
.el-carousel__item div {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 45px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #fce0ff;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #e4ffe0;
}

div#carousel{
  margin-top: -20px;
}

ul{
  text-align: left;
  list-style-type: none;
}
</style>
