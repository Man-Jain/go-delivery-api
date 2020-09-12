// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}


type tAuth struct {}
var Auth tAuth


func (_ tAuth) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Login", args).URL
}

func (_ tAuth) DeliveryLogin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.DeliveryLogin", args).URL
}

func (_ tAuth) RefreshToken(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.RefreshToken", args).URL
}


type tCookies struct {}
var Cookies tCookies


func (_ tCookies) GetCookies(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Cookies.GetCookies", args).URL
}

func (_ tCookies) CreateCookie(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Cookies.CreateCookie", args).URL
}

func (_ tCookies) GetCookie(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Cookies.GetCookie", args).URL
}


type tDelivery struct {}
var Delivery tDelivery


func (_ tDelivery) GetDeliveryPeople(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Delivery.GetDeliveryPeople", args).URL
}

func (_ tDelivery) GetDeliveryPerson(
		id uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Delivery.GetDeliveryPerson", args).URL
}

func (_ tDelivery) CompleteDelivery(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Delivery.CompleteDelivery", args).URL
}


type tOrders struct {}
var Orders tOrders


func (_ tOrders) GetOrders(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Orders.GetOrders", args).URL
}

func (_ tOrders) GetOrdersUser(
		userid uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userid", userid)
	return revel.MainRouter.Reverse("Orders.GetOrdersUser", args).URL
}

func (_ tOrders) GetOrdersDeliveryAgent(
		deliverypersonid uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "deliverypersonid", deliverypersonid)
	return revel.MainRouter.Reverse("Orders.GetOrdersDeliveryAgent", args).URL
}

func (_ tOrders) CreateOrder(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Orders.CreateOrder", args).URL
}

func (_ tOrders) GetOrder(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Orders.GetOrder", args).URL
}


type tUsers struct {}
var Users tUsers


func (_ tUsers) GetUsers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.GetUsers", args).URL
}

func (_ tUsers) GetUser(
		id uint,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Users.GetUser", args).URL
}

func (_ tUsers) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.Register", args).URL
}

func (_ tUsers) RegisterRoot(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.RegisterRoot", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


