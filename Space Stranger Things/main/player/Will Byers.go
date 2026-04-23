components {
  id: "player"
  component: "/main/player/player.script"
  position {
    x: 39.160717
    y: 21.636248
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"player\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    x: 16.0
    y: 16.0
  }
  scale {
    x: 0.054755
    y: 0.054755
    z: 0.054755
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
  "mask: \"chao\"\n"
  "mask: \"coin\"\n"
  "mask: \"enemy\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 15.0\n"
  "      y: 16.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 12.750838\n"
  "  data: 15.466666\n"
  "  data: 15.466666\n"
  "}\n"
  ""
}
