const { API } = require("../../utils/api");

Page({
  data: {
    folderList: [{name:'默认收藏'},{name:'金属'},{name:'人工智能'}],
    edit: false,
    fake_start: 0,
    fake: {
      name: '拖动假目标',
      top: 0
    },
    mvo: -1
  },
  jump: function(e) {
    if(!this.data.edit) {
      wx.navigateTo({
        url: '/pages/collect_paper_edit/index?name='+e.currentTarget.dataset.name,
      })
    }
  },
  delete: function(e) {
    let list = this.data.folderList;
    let item = list[e.currentTarget.dataset.index];
    list.splice(e.currentTarget.dataset.index,1);
    this.setData({
      folderList: list
    })
    wx.request({
      url: 'https://api.hust.online/doc-aid-test/folder/'+item.name,
      method: 'delete',
      success: function(rel) {
      },
      header: {
        'Authorization': 'Bearer '+ this.data.token
      }
    })
  },
  to_edit: function() {
    this.setData({
      edit: true
    })
  },
  to_save: function() {
    this.setData({
      edit: false
    })
    this.saveOrder()
  },
  mvs: function(e) {
    this.setData({
      mvo: e.currentTarget.dataset.index,
      y0: e.changedTouches[0].clientY,
      dy: e.changedTouches[0].clientY - e.currentTarget.offsetTop,
      fake_start: e.currentTarget.offsetTop - 86,
      fake: {
        name: e.currentTarget.dataset.name,
        top: e.currentTarget.offsetTop - 86
      }
    })
  },
  mv: function(e) {
    if(this.data.mvo > -1) {
      let folderList = this.data.folderList;
      let mvo = this.data.mvo;
      let temp = folderList[mvo];
      let fake = this.data.fake;
      fake.top = e.changedTouches[0].clientY - this.data.dy - 86;
      if(e.changedTouches[0].clientY - this.data.y0 > 25 && mvo < folderList.length - 1) {
        folderList[mvo] = folderList[mvo + 1];
        folderList[mvo + 1] = temp;
        mvo += 1;
        fake.name = folderList[mvo-1].name;
        this.setData({
          folderList,
          mvo,
          y0: e.changedTouches[0].clientY,
          dy: e.changedTouches[0].clientY - e.currentTarget.offsetTop
        })
      }else if(e.changedTouches[0].clientY - this.data.y0 < -25 && mvo > 0) {
        folderList[mvo] = folderList[mvo - 1];
        folderList[mvo - 1] = temp;
        mvo -= 1;
        fake.name = folderList[mvo+1].name;
        this.setData({
          folderList,
          mvo,
          y0: e.changedTouches[0].clientY,
          dy: e.changedTouches[0].clientY - e.currentTarget.offsetTop
        })
      }
      this.setData({
        fake
      })
    }
  },
  mve: function(e) {
    this.setData({
      mvo: -1
    })
  },
  saveOrder: function() {
    let orderList = this.data.folderList.map(item => item.name);
    orderList = ['默认收藏',...orderList];
    console.log(orderList);
    API.sortFavFolder(orderList).then(res => {
      console.log(res);
    })
  },
  getData() {
    let that = this
    API.getFav().then(res => {
      console.log(res);
      let folderList = res.data.data
            folderList = folderList.filter(item => item.name !== '默认收藏')
            that.setData({
              folderList
            })
    })

  },
  onLoad: function(e) {
    this.getData()
  }
})