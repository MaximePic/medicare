// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.queryAllMedic = function(){

		appFactory.queryAllMedic(function(data){
			console.log(data);
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});

			console.log(array);
			$scope.all_medic = array;
		});
	}

	$scope.queryMedic = function(){

		var id = $scope.medic_id;

		appFactory.queryMedic(id, function(data){
			$scope.query_medic = data;

			if ($scope.query_medic == "Could not locate medic"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordMedic = function(){

		appFactory.recordMedic($scope.medic, function(data){
			$scope.create_medic = data;
			$("#success_create").show();
		});
	}

	$scope.changeHolder = function(){

		appFactory.changeHolder($scope.holder, function(data){
			$scope.change_holder = data;
			if ($scope.change_holder == "Error: no tuna catch found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllMedic = function(callback){

    	$http.get('/get_all_medic/').success(function(output){
			callback(output)
		});
	}

	factory.queryMedic = function(id, callback){
    	$http.get('/get_medic/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordMedic = function(data, callback){

		data.location = data.longitude + ", "+ data.latitude;

		var medic = data.id + "-" + data.location;

    	$http.get('/add_medic/'+medic).success(function(output){
			callback(output)
		});
	}

	factory.changeHolder = function(data, callback){

		var holder = data.id + "-" + data.name;

    	$http.get('/change_holder/'+holder).success(function(output){
			callback(output)
		});
	}

	return factory;
});



