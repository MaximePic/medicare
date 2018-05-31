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

		var medic = data.id + "-" + data.refProduit + "-" + data.nomProduit 
		+ "-" + data.nomFabricant + "-" + data.numeroLot + "-" + data.dateExp 
		+ "-" + data.localisationVille + "-" + data.localisationEtablissement 
		+ "-" + data.dateCreation + "-" + data.dateReception + "-" + "false";



   	$http.get('/add_medic/'+ medic).success(function(output){
			callback(output)
		});
	}


	return factory;
});



