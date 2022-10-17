import React, { useState } from "react";
import styled from "styled-components";
import { P } from "../../interfaces/common";
import { randomNumber, randomString } from "../../utils/random";
import Checkbox from "../checkbox/Checkbox";
import Input from "../input/Input";
import T8y from "../t8y/T8y";

const Root = styled.div``;

const Table = styled.table`
    table-layout: fixed;
`;

interface SignupSuggestionProps extends P {}

const SignupSuggestion = (p: SignupSuggestionProps) => {
    // these random strings can be improved
    const [randomUsername] = useState<string>(getRandomUsername());
    const [randomPassword] = useState<string>(randomString(10));

    const [username, setUsername] = useState<string>(randomUsername);
    const [password, setPassword] = useState<string>(randomPassword);
    const [confirmPassword, setConfirmPassword] = useState<string>(randomPassword);
    const [passwordMatched, setPasswordMatched] = useState<boolean>(true);
    const [showPassword, setShowPassword] = useState<boolean>(false);

    return (
        <Root className={["flex flex-col ai-center", p.className].join(" ")}>
            <Table>
                <tbody>
                    <tr>
                        <td className="ta-end">
                            <T8y text="Username" />
                        </td>
                        <td className="ta-start">
                            <Input type="text" value={username} onChange={v => setUsername(v)} />
                        </td>
                    </tr>

                    <tr>
                        <td className="ta-end">
                            <T8y text="Password" />
                        </td>
                        <td className="ta-start">
                            <Input type="password" value={password} onChange={v => setPassword(v)} />
                        </td>
                    </tr>

                    <tr>
                        <td className="ta-end">
                            <T8y text="Confirm password" />
                        </td>
                        <td className="ta-start">
                            <Input type="password" value={confirmPassword} onChange={v => setConfirmPassword(v)} />
                        </td>
                    </tr>

                    <tr>
                        <td></td>
                        <td className="ta-start">
                            <Checkbox label="Show password" checked={showPassword} onChange={v => setShowPassword(v)} />
                        </td>
                    </tr>
                </tbody>
            </Table>
        </Root>
    );
};

export default SignupSuggestion;

const getRandomUsername = ():  string => {
    const name = animals[randomNumber(0, animals.length)].toLocaleLowerCase();
    const unique = randomNumber(1000, 9999);
    return name + unique;
};

const animals = [
    "Aardvark",
    "Albatross",
    "Alligator",
    "Alpaca",
    "Ant",
    "Anteater",
    "Antelope",
    "Ape",
    "Armadillo",
    "Donkey",
    "Baboon",
    "Badger",
    "Barracuda",
    "Bat",
    "Bear",
    "Beaver",
    "Bee",
    "Bison",
    "Boar",
    "Buffalo",
    "Butterfly",
    "Camel",
    "Capybara",
    "Caribou",
    "Cassowary",
    "Cat",
    "Caterpillar",
    "Cattle",
    "Chamois",
    "Cheetah",
    "Chicken",
    "Chimpanzee",
    "Chinchilla",
    "Chough",
    "Clam",
    "Cobra",
    "Cockroach",
    "Cod",
    "Cormorant",
    "Coyote",
    "Crab",
    "Crane",
    "Crocodile",
    "Crow",
    "Curlew",
    "Deer",
    "Dinosaur",
    "Dog",
    "Dogfish",
    "Dolphin",
    "Dotterel",
    "Dove",
    "Dragonfly",
    "Duck",
    "Dugong",
    "Dunlin",
    "Eagle",
    "Echidna",
    "Eel",
    "Eland",
    "Elephant",
    "Elk",
    "Emu",
    "Falcon",
    "Ferret",
    "Finch",
    "Fish",
    "Flamingo",
    "Fly",
    "Fox",
    "Frog",
    "Gaur",
    "Gazelle",
    "Gerbil",
    "Giraffe",
    "Gnat",
    "Gnu",
    "Goat",
    "Goldfinch",
    "Goldfish",
    "Goose",
    "Gorilla",
    "Goshawk",
    "Grasshopper",
    "Grouse",
    "Guanaco",
    "Gull",
    "Hamster",
    "Hare",
    "Hawk",
    "Hedgehog",
    "Heron",
    "Herring",
    "Hippopotamus",
    "Hornet",
    "Horse",
    "Human",
    "Hummingbird",
    "Hyena",
    "Ibex",
    "Ibis",
    "Jackal",
    "Jaguar",
    "Jay",
    "Jellyfish",
    "Kangaroo",
    "Kingfisher",
    "Koala",
    "Kookabura",
    "Kouprey",
    "Kudu",
    "Lapwing",
    "Lark",
    "Lemur",
    "Leopard",
    "Lion",
    "Llama",
    "Lobster",
    "Locust",
    "Loris",
    "Louse",
    "Lyrebird",
    "Magpie",
    "Mallard",
    "Manatee",
    "Mandrill",
    "Mantis",
    "Marten",
    "Meerkat",
    "Mink",
    "Mole",
    "Mongoose",
    "Monkey",
    "Moose",
    "Mosquito",
    "Mouse",
    "Mule",
    "Narwhal",
    "Newt",
    "Nightingale",
    "Octopus",
    "Okapi",
    "Opossum",
    "Oryx",
    "Ostrich",
    "Otter",
    "Owl",
    "Oyster",
    "Panther",
    "Parrot",
    "Partridge",
    "Peafowl",
    "Pelican",
    "Penguin",
    "Pheasant",
    "Pig",
    "Pigeon",
    "Pony",
    "Porcupine",
    "Porpoise",
    "Quail",
    "Quelea",
    "Quetzal",
    "Rabbit",
    "Raccoon",
    "Rail",
    "Ram",
    "Rat",
    "Raven",
    "Red deer",
    "Red panda",
    "Reindeer",
    "Rhinoceros",
    "Rook",
    "Salamander",
    "Salmon",
    "Sand Dollar",
    "Sandpiper",
    "Sardine",
    "Scorpion",
    "Seahorse",
    "Seal",
    "Shark",
    "Sheep",
    "Shrew",
    "Skunk",
    "Snail",
    "Snake",
    "Sparrow",
    "Spider",
    "Spoonbill",
    "Squid",
    "Squirrel",
    "Starling",
    "Stingray",
    "Stinkbug",
    "Stork",
    "Swallow",
    "Swan",
    "Tapir",
    "Tarsier",
    "Termite",
    "Tiger",
    "Toad",
    "Trout",
    "Turkey",
    "Turtle",
    "Viper",
    "Vulture",
    "Wallaby",
    "Walrus",
    "Wasp",
    "Weasel",
    "Whale",
    "Wildcat",
    "Wolf",
    "Wolverine",
    "Wombat",
    "Woodcock",
    "Woodpecker",
    "Worm",
    "Wren",
    "Yak",
    "Zebra"
];
