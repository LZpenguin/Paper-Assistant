// pages/add_keywords/index.js
Page({
  data: {
    keywordList: [{
      name: '经济全球化',
      add: false
    },{
      name: '经济全球',
      add: true
    },{
      name: '经济全球',
      add: true
    },{
      name: '经济全球',
      add: true
    },{
      name: '经济全球化',
      add: false
    },{
      name: '经济全球化',
      add: false
    }]
  },
  getData(key) {
    var that = this;
    wx.getStorage({
      key: 'token',
      success: res => {
        this.setData({
          'token': res.data
        })
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/keyword/search/'+key,
          success: (rel) => {
            let list = rel.data.data;
            that.setData({
              keywordList: list
            })
          },
          header: {
            'Authorization': 'Bearer '+ res.data
          }
        })
      }
    })
  },
  onLoad(e) {
    wx.getStorage({
      key: 'token',
      success: res => {
        this.setData({
          'token': res.data
        })
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/keyword',
          success: (rel) => {
            let list = rel.data.data;
            list = list.map(item => {
              return {
                sub: true,
                text: item
              }
            });
            if(list.length > 0) {
              this.setData({
                keywordList: list
              })
            }else {
              this.getData('文献')
            }
          },
          header: {
            'Authorization': 'Bearer '+ res.data
          }
        })
      }
    })
  },
  getValue(e) {
    this.getData(e.detail.value)
  },
  delete_key(e) {
    wx.request({
      url: 'https://api.hust.online/doc-aid-test/keyword',
      method: 'delete',
      data: [e.currentTarget.dataset.key],
      success: (rel) => {
        let list = this.data.keywordList;
        list.find(item => item.text == e.currentTarget.dataset.key).sub = false;
        this.setData({
          keywordList: list
        })
      },
      header: {
        'Authorization': 'Bearer '+ this.data.token
      }
    })
  },
  add_key(e) {
    wx.request({
      url: 'https://api.hust.online/doc-aid-test/keyword',
      method: 'post',
      data: [e.currentTarget.dataset.key],
      success: (rel) => {
        let list = this.data.keywordList;
        list.find(item => item.text == e.currentTarget.dataset.key).sub = true;
        this.setData({
          keywordList: list
        })
      },
      header: {
        'Authorization': 'Bearer '+ this.data.token
      }
    })
  }
})