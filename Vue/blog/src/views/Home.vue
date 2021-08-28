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
      style="height: 2000px;min-width: 1440px; background: rgba(20 ,20 ,20, 0); width: 100vw; position: absolute; top: 0; z-index: 20000; "
  >
  </vue-particles>
  <div id="main-menu">
    <Header> </Header>
    <div id="main-body" style="background-color: black">
      <div id="body-part1" style="min-width: 400px;" >
        <el-carousel indicator-position="outside" >
          <el-carousel-item >
            <img src="../assets/home/carousal1.png" alt="" style="max-height: 300px">
          </el-carousel-item>
          <el-carousel-item >
            <img src="../assets/home/carousal2.png" alt="" style="max-height: 300px">
          </el-carousel-item>
          <el-carousel-item >
            <img src="../assets/home/carousal3.png" alt="" style="max-height: 300px">
          </el-carousel-item>
          <el-carousel-item >
            <img src="../assets/home/carousal4.png" alt="" style="max-height: 300px">
          </el-carousel-item>
        </el-carousel>

        <div id="selector-div">
          <Selector> </Selector>
        </div>

        <suspense>
          <template #default>
            <HomeBlogPage />
          </template>
          <template #fallback>
            加载数据中....
          </template>
        </suspense>

      </div>
      <div id="body-part2" style="background-color: rgba(255, 255, 255, 0)">
        <div id="imgs" style="text-align: right; padding:0 10px 0 10px; min-width: 400px">
          <img src="../assets/home/Mypic.jpg" alt="" style="max-height: 300px; min-width: 100px; padding: 10px">
          <img src="../assets/home/Mypic2.jpg" alt="" style="max-height: 300px;  min-width: 100px; padding: 10px">
        </div>
          <div id="sticky-div">
              <div id="count-result" >
                <div class="item-count">
                    <document style="width: 50px; color: hotpink"/>
                  <div>
                    博客总数
                  </div>
                  <div>
                    {{blogsInfo.blogsNum}}
                  </div>
                </div>
                <div class="item-count">
                  <expand style="width: 50px; color: #3485ff" />
                  <div>
                    领域数目
                  </div>
                  <div>
                    {{blogsInfo.fieldsNum}}
                  </div>
                </div>
                <div class="item-count">
                  <files style="width: 50px; color: yellow" />
                  <div>
                    分类数目
                  </div>
                  <div>
                    {{ blogsInfo.classificationNum }}
                  </div>
                </div>
                <div class="item-count">
                  <document-checked style="width: 50px; color: greenyellow" />
                  <div>
                    标签数目
                  </div>
                  <div>
                    {{ blogsInfo.tagNum }}
                  </div>
                </div>
              </div>
            <div class="tools" >
              <span class="subtitle"> 个人简介 </span>
              <br>
              <ul class="tools-form">
                <span> 华中科技大学大二电信系学生。</span>
                <span style="display: block; font-weight: bold">开发经验:</span>
                <li class="contact-li">
                  计算机视觉一年转web前后端
                </li>
                <span style="display: block; font-weight: bold">技术栈:</span>
                <li class="contact-li">
                  主要学习使用Python、Golang; Java、C为辅
                </li>
              </ul>
            </div>
            <div id="contact-me" class="tools" style="position: sticky; bottom:0; z-index: 30000">
              <span class="subtitle"> 联系站长 </span>
              <ul class="tools-form">
                <li class="contact-li">
                  <img width="24" src="https://lsw-fast.lenovo.com.cn/appstore/normal/apps/2347-2020-05-21053954-1590053994120.png">
                  <span class="numbers">857998989</span>
                </li>
                <li class="contact-li">
                  <div style="background-color: white; display: inline; padding-top: 5px">
                    <svg height="24" aria-hidden="true" viewBox="0 0 16 16" version="1.1" width="24" data-view-component="true" class="octicon octicon-mark-github v-align-middle">
                      <path fill-rule="evenodd" d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"></path>
                    </svg>
                  </div>
                  <a class="numbers" href="https://github.com/XuYi-fei" style="position: relative; bottom: 5px; text-decoration: underline; cursor: pointer">https://github.com/XuYi-fei</a>
                </li>
                <li class="contact-li">
                  <img width="28" src="../assets/home/icon1.png" style="position: relative; bottom: 4px">
                  <a class="numbers" href="https://space.bilibili.com/25140014" style="position: relative; bottom: 5px; text-decoration: underline; cursor: pointer">https://space.bilibili.com/25140014(黑历史)</a>
                </li>
                <li class="contact-li"> 邮箱: <span class="numbers">857998989@qq.com / xuyifei@hust.edu.cn</span></li>
              </ul>
            </div>

            <div id="get-money" class="tools">
              <span class="subtitle"> 投喂食物 </span>
              <ul class="tools-form">
                <li class="pay">
                  <img height="50" src="../assets/home/pay1.png" class="logo">
                  <img height="100" src="../assets/home/pay3.png" class="logo">
                  <img height="50" src="../assets/home/pay2.png" class="logo">
                  <img height="100" src="../assets/home/pay4.png" class="logo">
                </li>
              </ul>
            </div>

            <div id="law" class="tools">
              <span>主办单位名称： 许逸飞</span>
              <span>主办单位性质： 个人</span>
              <span>网站备案/许可证号： <a class="law-a" href="https://icp.chinaz.com/xuyi-fei.xyz"> 鄂ICP备2021002829号-1 </a>
              </span>
              <span>审核时间： 2021-2-23</span>

            </div>

          </div>
        </div>
    </div>

  </div>
</template>

<script>
// @ is an alias to /src
import Header from "../components/Header";
import Selector from "../components/selections/Selector";
import Footer from "../components/Footer";
import { Expand, Document, Files, DocumentChecked } from '@element-plus/icons'
import HomeBlogPage from "../components/home/HomeBlogPage";
import {getBlogsInfo} from "../api/home";
export default {
  name: 'Home',
  components: {
    Footer,
    Selector,
    Header,
    Expand,
    Files,
    Document,
    DocumentChecked,
    HomeBlogPage
  },
  async created(){
    const response = await getBlogsInfo()
    this.blogsInfo = response.data
  },
  data(){
    return {
      blogs: {},
      imgPaths: [
        '../assets/home/carousal1.png',
        '../assets/home/carousal2.png',
        '../assets/home/carousal3.png',
        '../assets/home/carousal4.png',
      ],
      blogsInfo: {},
      dialogImageUrl: '',
      dialogVisible: false,
      style: {
        "background-color": "rgb(20 ,20 ,20)",
        height: "2000px"
      }
    }
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePictureCardPreview(file) {
      console.log(file)
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    }
  }
}
</script>

<style scoped>
/* 组件排版 */

/* 轮播图排版 */
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}

.el-carousel__item{
  background: transparent;
}


a{
  text-decoration: none;
  color: #3485ff;
  font-family: "STKaiti", sans-serif;
  position: relative;
  z-index: 30000;
}

a.law-a{
  color: rgb(220, 220, 220);
  text-decoration: underline;
}
/* 主体排版 */
div#main-menu{
  position: absolute;
  top: 0;
  overflow: auto;
  min-width: 1440px;
  width: 100vw;
  background: transparent;
  display: flex;
  flex-direction: column;
}

/* 左右部分的布局 */
div#main-body{
  display: flex;
  min-width: 1440px;
  overflow: auto;
  flex-direction: row;
}

div#body-part1{
  flex: 72%;
}

div#body-part2{
  flex: 28%;

}

/* 选择器定位 */
div#selector-div{
  display: flex;
  justify-content: left;
  padding-left: 25px;
}


div#imgs{
  display: flex;
  justify-content: space-evenly;
}

div#count-result{
  display: flex;
  color: white;
  justify-content: space-evenly;
}

div.tools{
  border-color: #99a9bf;
  border-style: solid;
  display: flex;
  flex-direction: column;
  padding: 10px;
  margin: 5px;
}

ul.tools-form{
  padding: 5px;
  text-align: left;
  font-size: 18px;
  list-style: none;
  color: white;
}

li.contact-li{
  padding: 5px;
  position: relative;
  z-index: 30000;
}

.numbers{
  padding-left: 5px;
  position: relative;
}

div#sticky-div{
  display: flex;
  flex-direction: column;

}

li.pay{
  display: flex;
  justify-content: space-evenly;
}

span.subtitle{
  text-align: left;
  font-size: large;
  font-family: 'Microsoft YaHei UI', sans-serif;
  color: #9dff7c
}

div#law{
  color: #f9ffed;
  text-align: left;
  font-size: large;
}
</style>
