let {API} = require('../../utils/api')
Page({
  data: {
    paper: {
      "id": "a49a9762-3790-4b4f-adbf-4577a35b1df7",
      "title": "应对气候变化的农业经济研究前沿与政策实践政策实践政策实践政策实践7",
      "authors": ["陈朴", "林垚", "刘凯"],
      "magazine": "经济研究",
      "date": "2021-06-20",
      "url": "https://kns.cnki.net/...",
      "year": 2021,
      "volume": 56,
      "issue": 6,
      "intro": "本文通过构建...",
      "fav": true,
      "keywords": ["统一大市场", "区域贸易", "劳动力流动", "经济增长"],
      "suggest": [
        {
          "id": "a49a9762-3790-4b4f-adbf-4577a35b1df7",
          "title": "应对气候变化的农业经济研究前沿与政策实践政策实践政策实践政策实践",
          "authors": ["陈朴", "林垚", "刘凯"],
          "magazine": "经济研究",
          "year": 2021,
          "issue": 3
        },
        {
          "id": "a49a9762-3790-4b4f-adbf-4577a35b1df7",
          "title": "应对气候变化的农业经济研究前沿与政策实践政策实践政策实践政策实践",
          "authors": ["陈朴", "林垚", "刘凯"],
          "magazine": "经济研究",
          "year": 2021,
          "issue": 3
        },
        {
          "id": "a49a9762-3790-4b4f-adbf-4577a35b1df7",
          "title": "应对气候变化的农业经济研究前沿与政策实践政策实践政策实践政策实践",
          "authors": ["陈朴", "林垚", "刘凯"],
          "magazine": "经济研究",
          "year": 2021,
          "issue": 3
        }
      ]
    },
    btn_collect: [600,985],
    btn_share: [600,1125],
    to_collect: false,
    to_add_folder: false,
    showRank: false
  },
  //弹出新建文件夹窗口
  to_add_folder: function(e) {
    this.setData({
      to_add_folder: true
    })
  },
  //关闭新建文件夹窗口
  cancel_add_folder(e) {
    this.setData({
      to_add_folder: false
    })
  },
  //新建文件夹名数据双向绑定
  bindInput(e) {
    this.setData({
      input: e.detail.value
    })
  },
  //新建文件夹
  new_folder() {
    wx.request({
      url: 'https://api.hust.online/doc-aid-test/folder/'+this.data.input,
      method: 'post',
      success: (rel) => {
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/fav',
          success: (rel) =>{
            this.setData({
              folderList: rel.data.data
            })
          },
          header: {
            'Authorization': 'Bearer '+ this.data.token
          }
        })
        this.setData({
          input: '',
          to_add_folder: false
        })
      },
      header: {
        'Authorization': 'Bearer '+ this.data.token
      }
    })
  },
  //去选择收藏夹添加收藏
  add_collect: function(e) {
    this.setData({
      to_collect: true
    })
  },
  //关闭选择收藏夹窗口
  cancel_add_collect: function() {
    this.setData({
      to_collect: false
    })
  },
  //添加收藏
  collect_to: function(e) {
    let name = e.currentTarget.dataset.name;
    let id = this.data.paper.id;
    API.addFav(name,[id]).then(res => {
      console.log(res);
      let paper = this.data.paper;
        paper.fav = true;
        this.setData({
          paper,
          to_collect: false
        })
        wx.setStorage({
          data: paper.id,
          key: 'collect_in_detail',
        })
        wx.showToast({
          title: '收藏成功',
          icon: 'success',
          duration: 500
        })
    })
  },
  //取消收藏
  cancel_collect: function(e) {
    let id = e.currentTarget.dataset.id;
    let folderList = this.data.folderList;
    let name = folderList.find(item => {
      return item.content.find(item2 => item2.id == id)
    }).name
    API.delFav(name,[id]).then(res => {
      console.log(res);
      let paper = this.data.paper;
      paper.fav = false;
      this.setData({
        paper
      })
      wx.setStorage({
        data: paper.id,
        key: 'cancel_collect_in_detail',
      })
    })
  },
  //分享
  onShareAppMessage:function(from,target,webViewUrl)	{
    wx.getStorage({
      key: 'shared',
      success: res => {
        let shared = res.data
        let shared_same_index = shared.findIndex(item => item.id == this.data.paper.id)
        if(shared_same_index>=0) {
          shared.splice(shared_same_index,1)
        }
        shared.unshift(this.data.paper)
        if(shared.length>50) {
          shared = shared.slice(0,50)
        }
        wx.setStorage({
          data: shared,
          key: 'shared',
        })
      },
      fail: err => {
        let shared = [this.data.paper]
        wx.setStorage({
          data: shared,
          key: 'shared',
        })
      }
    })
	  return{
	    title: this.data.paper.title,
	    path: "/pages/detail/index?id="+this.data.paper.id,
		  imageUrl:''
	  }
  },
  //订阅期刊
  add_sub: function(e) {
    let name = e.currentTarget.dataset.name
    let that = this
    API.sub(name).then(res => {
      let paper = that.data.paper;
      paper.magazineSub = true;
      that.setData({
        paper
      })
    })
  },
  //取消订阅期刊
  delete_sub: function(e) {
    let names = [e.currentTarget.dataset.name]
    let that = this
    API.unsub(names).then(res => {
      let paper = that.data.paper;
      paper.magazineSub = false;
      that.setData({
        paper
      })
    })
  },
  //复制DOI
  copy_doi: function() {
    wx.setClipboardData({
      data: this.data.paper.title
    })
  },
  //跳转到期刊
  link_to: function(e) {
    let base_url = 'https://s.wanfangdata.com.cn/periodical?q='
    wx.setClipboardData({
      data: base_url + this.data.paper.title
    })
    // wx.navigateTo({
    //   url: '/pages/outer/index?link='+e.currentTarget.dataset.link,
    // })
  },
  //浮动按钮归右
  to_right: function(e) {
    let ratio = this.data.ratio;
    let height = this.data.height;
    let x = e.currentTarget.dataset.name;
    let y = (e.changedTouches[0].clientY)*ratio;
    let y_ = x==0?this.data.btn_share[1]:this.data.btn_collect[1];
    if(y>=y_&&y-y_<140) {
       y = y_ + 140;
    }else if (y<y_&&y_-y<140) {
      y= y_ - 140;
    }
    this.setData({
      [x==0?'btn_collect':'btn_share']: [600,y]
    })
  },
  //跳转期刊详情页
  jump_periodical_detail(e) {
    wx.navigateTo({
      url: '/pages/periodical_detail/index?name='+e.currentTarget.dataset.magazine,
    })
  },
  //跳转文章详情页
  jump_detail(e) {
    wx.navigateTo({
      url: '/pages/detail/index?id='+e.currentTarget.dataset.id,
    })
  },
  //获取页面数据
  getData(id) {
    var that = this;
    API.getFav().then(res => {
      let folderList = res.data.data;
      that.setData({
        folderList
      })
    })
    API.detail(id).then(res => {
      let paper = res.data.data;
      if(paper.magazine.length > 10) {
        paper['magazine_fixed'] = paper.magazine.slice(0,10) + '...'
      }else {
        paper['magazine_fixed'] = paper.magazine
      }
      paper.authors = paper.authors.slice(0,4);
      paper.suggest&&paper.suggest.forEach(item => {
        item.authors = item.authors.slice(0,4);
      })
      that.setData({
        paper
      })
    })
  },
  onLoad: function (e) {
    this.setData({
      id: e.id
    })
    this.getData(e.id)
    wx.getSystemInfo({
      success: (res) => {
        // 获取可使用窗口宽度
        let clientHeight = res.windowHeight;
        // 获取可使用窗口高度
        let clientWidth = res.windowWidth;
        // 算出比例
        let ratio = 750 / clientWidth;
        // 算出高度(单位rpx)
        let height = clientHeight * ratio;
        // 设置高度
        this.setData({
          height: height,
          ratio: ratio
        });
      }
    });
  }
})