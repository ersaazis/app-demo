<?php

/** @var \Laravel\Lumen\Routing\Router $router */

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It is a breeze. Simply tell Lumen the URIs it should respond to
| and give it the Closure to call when that URI is requested.
|
*/

$router->get('/', function () use ($router) {
    return $router->app->version();
});

$router->get('/health', function () {
    return response('OK', 200);
});

$router->get('/go', function () {
    sleep(5);
    return response('OK after 5s', 200);
});

$router->get('/query', function () {
    return response('Database connection error (simulated)', 500);
});
