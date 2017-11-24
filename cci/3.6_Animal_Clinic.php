<?php

// 3.6 Animal Shelter
// An animal shelter, which holds only dogs and cats, operates on a strictly "first in, first out" basis.
// People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
// or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
// that type).
// They cannot select which specific animal they would like.
//
// Create the data structures to maintain this system and implement operations such as
// enqueue, dequeueAny, dequeueDog and dequeueCat.
//
// You may use the built-in Linked list data structure.
/**

SOLUTION
...use separate queues for dogs and cats, and to place them within a wrapper class called AnimalQueue.
We then store some sort of timestamp to mark when each animal was enqueued.
When we call dequeueAny, we peek at the heads of both the dog and cat queue and return the oldest.

 */
abstract class Animal
{

    protected $name = null;
    protected $order = null;

    public function Animal($n)
    {
        $this->name = $n;
    }

    public function setorder($ord)
    {
        $this->order = $ord;
    }

    public function getOrder()
    {
        return $this->order;
    }

    public function isOlderThan(Animal $a)
    {
        return $this->order < $a->getOrder();
    }
}

class AnimalQueue
{
    public $dogs = [];
    public $cats = [];
    private $order = 0;

    public function enqueue(Animal $a)
    {
        $a->setorder($this->order);
        $this->order++;

        if ($a instanceof Dog) {
            $this->dogs[] = $a;
        } else if ($a instanceof Cat) {
            $this->cats[] = $a;
        }
    }

    public function dequeueAny()
    {
        if (sizeof($this->dogs) === 0) {
            return $this->dequeueCats();
        } elseif (sizeof($this->cats) === 0) {
            return $this->dequeueDogs();
        }

        $dog = end($this->dogs);
        $cat = end($this->cats);

        if ($dog->isOlderThan($cat)) {
            return $this->dequeueDogs();
        } else {
            return $this->dequeueCats();
        }
    }

    public function dequeueDogs()
    {
        return array_shift($this->dogs);
    }

    public function dequeueCats()
    {
        return array_shift($this->cats);
    }
}

class Dog extends Animal
{
}

class Cat extends Animal
{
}

$dog0 = new Dog("Norton");
$dog1 = new Dog("Memphis");
$cat0 = new Cat("Musia");
$cat1 = new Cat("Jessie");

$clinic = new AnimalQueue();


$clinic->enqueue($dog0);
$clinic->enqueue($cat0);
$clinic->enqueue($cat1);
$clinic->enqueue($dog1);

$someAnimal = $clinic->dequeueAny();
var_dump($someAnimal);

$d = $clinic->dequeueDogs();
var_dump($d);

$d = $clinic->dequeueDogs();
var_dump($d);

$d = $clinic->dequeueDogs();
var_dump($d);

$c = $clinic->dequeueCats();
var_dump($c);