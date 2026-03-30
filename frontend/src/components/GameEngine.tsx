import { useEffect, useRef } from "react";
import Phaser from "phaser";


export default function Engine({peerConnection} : {peerConnection : RTCPeerConnection}) {
    const game = useRef<Phaser.Game>;
    useEffect(() => {
        game.current = new Phaser.Game({
            type: Phaser.AUTO,
            width: 800,
            height: 600,
            backgroundColor: '#000000',
            scene: {
                preload: preload,
                create: create,
                update: update
            },
            parent: 'phaser-game',
            physics: {
                default: 'arcade',
                arcade: {
                    gravity: {x:0, y: 0}, 
                    debug: true,
                },
            }
        });
    }, []);

    function preload() {
        // load assets
        // e.g. this.load.image(name, path)
    }

    function create() {
        // create env / display
        game.current.up = game.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.W);
    }

    function update(time : number, delta: number) {
        // get user input
        // send to server
    }

    return <div id="phaser-game"></div>;
}