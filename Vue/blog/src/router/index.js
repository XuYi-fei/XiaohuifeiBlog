import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import PersonalSpace from "../components/Personal/PersonalSpace";
import PersonalInfo from "../components/Personal/PersonalInfo";
import Blog from "../views/Blog";
import CreateBlog from "../components/Personal/CreateBlog";
import EditBlog from "../components/Personal/EditBlog";
import PersonalDetail from "../components/Personal/PersonalDetail";
import AllBlogs from "../views/AllBlogs";
import Sentences from "../views/Sentences";
import Classification from "../views/Classification";
import UpdateBlog from "../components/Personal/UpdateBlog";
import Navgation from "../components/Personal/Navgation";
import UpdateBread from "../components/Personal/UpdateBread";
import RootBread from "../components/Personal/RootBread";
import Admin from "../views/Admin";
import AdminIndex from "../components/admin/AdminIndex";
import AdminComplaints from "../components/admin/AdminComplaints";
import AdminBlogs from "../components/admin/AdminBlogs";
import AdminSentences from "../components/admin/AdminSentences";
import AdminCheckBlog from "../components/admin/AdminCheckBlog";
import PersonalSentences from "../components/Personal/PersonalSentences";
import Inbox from "../components/Personal/Inbox";
import AdminBlogsAll from "../components/admin/AdminBlogsAll";
import AdminSentencesAll from "../components/admin/AdminSentencesAll";
import ForgettingPassword from "../components/password/ForgettingPassword";
import InboxAllPage from "../components/Personal/InboxAllPage";
import inboxMessagePage from "../components/Personal/inboxMessagePage";
import Comments from "../views/Comments";

const routes = [
  {
    path: '/',
    redirect: 'home'
  },
  {
    path: '/home',
    name: 'Home',
    components: {
      home: Home
    }
  },
  {
    path: '/allBlogs',
    name: 'allBlogs',
    components: {
      home: AllBlogs,
    }
  },
  {
    path: '/sentences',
    name: 'sentences',
    components: {
      home: Sentences,
    }
  },
  {
    path: '/comments',
    name: 'comments',
    components: {
      home: Comments,
    }
  },
  {
    path: '/classification',
    redirect: '/classification/recentBlogs/latest'
  },
  {
    path: '/classification/:field([^/]+)/:classes([^/]+)',
    name: 'classification',
    components: {
      home: Classification,
    },
  },
  {
    path: '/admin',
    name: 'admin',
    redirect: '/admin/index',
    components: {
      home: Admin
    },
    children:[
      {
        path:'/',
        redirect: '/index'
      },
      {
        name: 'checkblogs',
        path: 'checkblogs/:blogId',
        components: {
          body: AdminCheckBlog
        }
      },
      {
        path: 'index',
        name: 'index',
        components:{
          indexView: AdminIndex,
        }
      },
      {
        path: 'sentence',
        name: 'sentence',
        components:{
          sentencesView: AdminSentences,
        }
      },
      {
        path: 'blog',
        name: 'blog',
        components:{
          blogsView: AdminBlogs,
        }
      },
      {
        path: 'blogAll',
        name: 'blogAll',
        components:{
          blogsAllView: AdminBlogsAll,
        }
      },
      {
        path: 'sentenceAll',
        name: 'sentenceAll',
        components:{
          sentencesAllView: AdminSentencesAll,
        }
      },
      {
        path: 'complaint',
        name: 'complaint',
        components:{
          complaintsView: AdminComplaints,
        }
      }
    ]
  },
  {
    name: 'forgettingPassword',
    path: '/forgettingPassword',
    components:{
      home: ForgettingPassword
    },
  },
  {
    path: '/personal/:userId([^/]+)/',
    name: 'personal',
    components: {
      personal: PersonalSpace
    },
    children: [
      {
        name: 'personal-info',
        path: 'personal-info',
        components: {
          assets: PersonalInfo,
          nav: Navgation,
          bread: RootBread
        }
      },
      {
        name: 'personal-detail',
        path: 'personal-detail',
        components: {
          assets: PersonalDetail,
          nav: Navgation,
          bread: RootBread

        }
      },
      {
        name: 'createblog',
        path: 'createblog',
        components:{
          assets: CreateBlog,
          nav: Navgation,
          bread: RootBread

        }
      },
      {
        name: 'editblog',
        path: 'editblog',
        components:{
          assets: EditBlog,
          nav: Navgation,
          bread: RootBread
        },
      },
      {
        name: 'inbox',
        path: 'inbox',
        components:{
          assets: Inbox,
          nav: Navgation,
          bread: RootBread
        },
        children:[
          {
            path: 'inboxAll',
            name: 'inboxAll',
            inboxHome: InboxAllPage
          },
          {
            path: 'inboxMessage/:messageId([^/]+)',
            name: 'inboxMessage',
            inboxHome: inboxMessagePage
          }
        ]
      },

      {
        name: 'editSentence',
        path: 'editSentence',
        components:{
          assets: PersonalSentences,
          nav: Navgation,
          bread: RootBread
        },
      },
      {
        name: 'updateblog',
        path: 'editblog/update/:blogId',
        components:{
          assets: UpdateBlog,
          bread: UpdateBread
        },
      },

    ],
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/blogs/:userId([^/]+)/article/details/:blogId',
    name: 'Blog',
    components: {
      blog: Blog
    },
  },

]

const router = createRouter({
  mode: 'history',
  base: '/root/Blog',
  history: createWebHistory("/"),
  routes
})

export default router
