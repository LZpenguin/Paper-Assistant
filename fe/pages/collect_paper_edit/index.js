const { API } = require("../../utils/api");

Page({
  data: {
    folderList: [],
    paperList: [],
    edit: false,
    count:0,
    to_move: false,
    to_add_folder: false
  },
  cancel: function() {
    let paperList = this.data.paperList;
    paperList.forEach(item => {
      item.choosen = false
    });
    this.setData({
      paperList,
      edit: false,
      count: 0
    })
  },
  all: function() {
    let paperList = this.data.paperList;
    let count = 0;
    paperList.forEach(item => {
      item.choosen = true
      count++
    });
    this.setData({
      paperList,
      edit: true,
      count
    })
  },
  dotTap: function(e) {
    let edit = false;
    let count = this.data.count;
    let paperList = this.data.paperList;
    let item = paperList.find(item => item.id == e.currentTarget.dataset.id);
    if(item.choosen) {
      item.choosen = false;
      count--;
    }else {
      item.choosen = true;
      count++;
    }
    if(paperList.find(item => item.choosen == true)) {
      edit = true;
    }
    this.setData({
      paperList,
      edit,
      count
    })
  },
  delete() {
    let paperList = this.data.paperList;
    let delList = [];
    paperList.forEach(item => {
      if(item.choosen) {
        delList.push(item.id)
      }
    });
    API.delFav(this.data.folderName,delList).then(res => {
      if(this.data.jump_back) {
        wx.navigateBack()
      }else {
        wx.redirectTo({
          url: '/pages/collect_paper_edit/index?name='+this.data.folderName,
        })
      }
    })
  },
  move() {
    if(this.data.count > 0) {
      this.setData({
        to_move: true
      })
      let moveList = [];
      let paperList = this.data.paperList;
      paperList.forEach(item => {
        if(item.choosen) {
          moveList.push(item.id)
        }
      })
    }
  },
  cancel_move() {
    this.setData({
      to_move: false
    })
  },
  move_to(e) {
    let paperList = this.data.paperList;
    let delList = [];
    let name = 
    paperList.forEach(item => {
      if(item.choosen) {
        delList.push(item.id)
      }
    });
    let p1 = API.delFav(this.data.folderName,delList)
    let p2 = API.addFav(e.currentTarget.dataset.name,delList)
    Promise.all(p1,p2).then(res => {
      if(this.data.jump_back) {
        wx.navigateBack()
      }else {
        wx.redirectTo({
          url: '/pages/collect_paper_edit/index?name='+this.data.folderName,
        })
      }
    })
  },
  getData(cb) {
    var that = this;
    API.getFav().then(res => {
      let list = res.data.data;
      that.setData({
        folderList: list,
      })
      if(cb) {
        cb(list)
      }
    })
  },
  onLoad(e) {
    this.getData((folderList) => {
      let item = folderList.find(item => item.name == e.name);
      if(!item) {
        wx.redirectTo({
          url: '/pages/collect_edit/index',
        })
      }else {
        let paperList = item.content;
        if(e.id) {
          paperList.map(item => {
            item['choosen']=false
            if(item.id == e.id) {
              item['choosen']=true
            }
            item.authors = item.authors.slice(0,4)
          })
          this.setData({
            jump_back: true,
            paperList,
            edit: true,
            count: 1
          })
        }else {
          paperList.map(item => item['choosen']=false)
          this.setData({
            paperList
          })
        }
      }
    });
    wx.setNavigationBarTitle({
      title: e.name,
    })
    this.setData({
      folderName: e.name
    })
    wx.getStorage({
      key: 'userInfo',
      success: (res) => {
        this.setData({
          nickName: res.data.nickName
        })
      }
    })
  },

  // add_folder
  add_folder(e) {
    this.setData({
      to_add_folder: true
    })
  },
  cancel_add_folder(e) {
    this.setData({
      to_add_folder: false
    })
  },
  bindInput(e) {
    this.setData({
      input: e.detail.value
    })
  },
  new_folder() {
    API.addFavFolder(this.data.input).then(res => {
      this.getData()
      this.setData({
        input: '',
        to_add_folder: false
      })
    })
  },
  getUserProfile() {
    wx.getUserProfile({
      desc: '用户登录',
      success: (res) => {
        this.setData({
          userInfo: res.userInfo
        })
        wx.setStorage({
          data: res.userInfo,
          key: 'userInfo',
        })
      }
    })
  },
  // share
  onShareAppMessage() {
    let shareList = [];
    let idList = [];
    let ids = '';
    this.data.paperList.forEach(item => {
      if(item.choosen) {
        idList.push(item.id)
        shareList.push(item)
      }
    })
    wx.getStorage({
      key: 'shared',
      success: res => {
        let shared = res.data
        shared = shared.filter(item => {
          return idList.every(id => id != item.id)
        })
        shared = [...shareList,...shared];
        if(shared.length>50) {
          shared = shared.slice(0,50)
        }
        wx.setStorage({
          data: shared,
          key: 'shared',
        })
      },
      fail: err => {
        let shared = [shareList]
        wx.setStorage({
          data: shared,
          key: 'shared',
        })
      }
    })
    ids = idList.join('+');
    return {
      title: '来自'+this.data.nickName+'的文献分享',
      path: '/pages/receive/index?nickName='+ this.data.nickName +'&ids='+ids
    }
  }
})