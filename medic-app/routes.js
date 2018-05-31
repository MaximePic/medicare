//SPDX-License-Identifier: Apache-2.0

var controller = require('./controller.js');

module.exports = function(app){

  app.get('/get_medic/:id', function(req, res){
    controller.get_medic(req, res);
  });
  app.get('/add_medic/:medic', function(req, res){
    controller.add_medic(req, res);
  });
  app.get('/get_all_medic', function(req, res){
    controller.get_all_medic(req, res);
  });
}

