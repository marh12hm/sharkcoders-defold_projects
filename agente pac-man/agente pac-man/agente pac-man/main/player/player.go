components {
  id: "player"
  component: "/main/scripts/player.script"
}
components {
  id: "morreste1"
  component: "/main/sounds/morreste.sound"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Player\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/main.atlas\"\n"
  "}\n"
  ""
  position {
    x: -1.0
  }
  scale {
    x: 0.84
    y: 0.84
    z: 0.84
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"door\"\n"
  "mask: \"documento\"\n"
  "mask: \"wall\"\n"
  "mask: \"vision\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: -3.0\n"
  "      y: -12.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"sphere.\"\n"
  "  }\n"
  "  data: 34.633972\n"
  "}\n"
  ""
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/main/sounds/freesound_community-blue-blood-ogg-67848.ogg\"\n"
  "group: \"sound\"\n"
  ""
}
