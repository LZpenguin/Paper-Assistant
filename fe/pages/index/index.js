let {API} = require('../../utils/api')
Page({
  data: {
    paperList: [],
    loading: false,
    collect: false,
    waiting: false
  },
  //获取本页数据
  getData(cb) {
    API.document().then(res => {
      this.setData({
        paperList: res.data.data
      })
      cb && cb(res.data.data)
    })
  },
  //添加收藏
  add_collect: function(e) {
    let id = e.currentTarget.dataset.id;
    let paperList = this.data.paperList;
    let index = paperList.findIndex(item => item.id == id);
    fav_add('默认收藏',id,this.data.token,res => {
      paperList[index].fav = true;
      this.setData({
        paperList
      })
    })
    
  },
  //取消收藏
  cancel_collect: function(e) {
    let id = e.currentTarget.dataset.id;
    let paperList = this.data.paperList;
    let index = paperList.findIndex(item => item.id == id);
    fav_del(id,this.data.token,res => {
      paperList[index].fav = false;
      this.setData({
        paperList
      })
    })
  },
  //跳转到文章详情页
  jump_detail(e) {
    wx.navigateTo({
      url: '/pages/detail/index?id='+e.currentTarget.dataset.id,
    })
  },
  //跳转到期刊详情页
  jump_periodical_detail(e) {
    wx.navigateTo({
      url: '/pages/periodical_detail/index?name='+e.currentTarget.dataset.magazine,
    })
  },
  //跳转到期刊导航页
  jump_periodical_navi(e) {
    wx.navigateTo({
      url: '/pages/periodical_navigator/index',
    })
    wx.setStorage({
      key: 'fromIndex',
      data: true
    })
  },
  //下拉刷新
  onPullDownRefresh(e) {
    this.getData((paperList) => {
      paperList = paperList.map(item => {
        item.authors = item.authors.slice(0,4);
        return item;
      })
      this.setData({
        paperList
      })
      setTimeout(wx.stopPullDownRefresh,250)
    })
  },
  //上滑刷新
  onReachBottom(e) {
    this.setData({
      loading: true
    })
    this.getData((paperList) => {
      setTimeout(() => {
        let newList = this.data.paperList.concat(paperList)
        newList = e=newList.map(item => {
          item.authors = item.authors.slice(0,4);
          return item;
        })
       this.setData({
          paperList: newList,
          loading: false
        })
      }, 1000);
    })
  },

  // 获取用户信息
  getUserProfile() {
    if(this.data.userInfo) {
      return
    }
    wx.getUserProfile({
      desc: '用户登录',
      success: res => {
        this.setData({
          userInfo: res.userInfo
        })
        wx.setStorageSync('userInfo', res.userInfo)
      }
    })
  },

  onLoad: function() {
    this.getData()
  },
  onShow() {
    wx.getStorage({
      key: 'collect_in_detail',
      success: rel => {
        let paperList = this.data.paperList
        paperList.find(item => item.id == rel.data).fav = true
        this.setData({
          paperList
        })
        wx.removeStorage({
          key: 'collect_in_detail',
        })
      }
    })
    wx.getStorage({
      key: 'cancel_collect_in_detail',
      success: rel => {
        let paperList = this.data.paperList
        paperList.find(item => item.id == rel.data).fav = false
        this.setData({
          paperList
        })
        wx.removeStorage({
          key: 'cancel_collect_in_detail',
        })
      }
    })
    wx.getStorage({
      key: 'fromIndex',
      success: res => {
        if(res.data) {
          wx.switchTab({
            url: '/pages/user/index',
          })
        }
      }
    })
  }
})
